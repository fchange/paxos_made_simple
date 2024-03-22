package model

import (
	"fmt"
	"time"
)

var electionInterval = 5 * time.Second // 选举间隔为5秒

// Proposals 是一个映射（map），它记录了所有已提出的提案，其中键是提案编号，值是指向Proposal结构体的指针。
// Promises 是一个布尔映射，它记录了节点已经做出的所有承诺的编号。
// State 是NodeState结构体，它包含了节点的当前状态信息，例如已收到的最高编号的承诺和已选定的值（如果有）。
// Node 表示参与Paxos算法的节点。
type Node struct {
	ID      int
	Env     *Environment
	State   NodeState    // 节点的状态
	Channel chan Message // 用于接收消息的channel
	// 其他字段...
}

func NewNode(ID int, env *Environment) *Node {
	return &Node{
		ID:  ID,
		Env: env,
		State: NodeState{
			ProposalNumber:    0,
			ProposalValue:     nil,
			Leader:            0,
			ElectionResults:   make(map[int]bool),
			LastHeartBeatTime: time.Time{},
		},
		Channel: make(chan Message, 1024),
	}
}

// Background 方法启动看门狗定时器，用于检测领导者是否离线，并在必要时触发选举。
func (node *Node) Background() {
	heartbeatInterval := 3 * time.Second // 假设心跳间隔为3秒
	ticker := time.NewTicker(heartbeatInterval)

	go func() {
		for range ticker.C {
			node.Channel <- Message{Type: MessageTypeMonitorLeader}
		}
	}()
}

func (node *Node) ParticipateInPaxos() {
	node.Background()

	fmt.Println("Node", node.ID, "is participating in Paxos")
	for {
		select {
		case msg := <-node.Channel:
			// 处理来自其他节点的消息
			fmt.Println("Node", node.ID, "received message:", msg.String())

			switch msg.Type {
			case MessageTypeMonitorLeader:
				node.MonitorLeader()
			case MessageTypeHeartBeat:
				node.HandleHeartBeat(msg)
			case MessageTypePrepare:
				node.HandlePrepare(msg)
			case MessageTypePromise:
				node.HandlePromise(msg)
			case MessageTypeAccept:
				node.HandleAccept(msg)
			case MessageTypeAccepted:
				node.HandleAccepted(msg)
			default:
				// handle default case if needed
				fmt.Println("Unknown message type:", msg.Type)
			}
		}
	}
}

// MonitorLeader 检查是否需要开始新的选举。
func (node *Node) MonitorLeader() {
	currentTime := time.Now()
	if currentTime.Sub(node.State.LastHeartBeatTime) > electionInterval {
		// 如果自上次接收到领导者消息以来已经超过选举间隔，则认为领导者已离线
		node.StartElection()
	}
}

func (node *Node) HandleHeartBeat(msg Message) {
	node.State.LastHeartBeatTime = time.Now()
}

// HandlePrepare 处理来自Proposer的Prepare请求。
func (node *Node) HandlePrepare(msg Message) {
	// 检查是否需要响应Prepare请求
	if msg.Number > node.State.ProposalNumber {
		node.State.ProposalNumber = msg.Number
		response := Message{
			Type:   MessageTypePromise,
			Number: msg.Number,
			From:   node.ID,
		}
		// 发送Promise响应给Proposer
		node.Env.ReceiveMessage(msg.From, response)
	}
}

// HandleAccept 处理来自Proposer的Accept请求。
func (node *Node) HandleAccept(msg Message) {
	// 检查Accept请求是否符合之前的Promise
	if msg.Number == node.State.ProposalNumber {
		node.State.ProposalNumber = msg.Number // 更新提案编号
		node.State.ProposalValue = msg.Value   // 更新提案值
		response := Message{
			Type:   MessageTypeAccepted,
			Number: msg.Number,
			Value:  msg.Value,
			From:   node.ID,
		}
		// 发送Accepted响应给Proposer
		node.Env.ReceiveMessage(msg.From, response)
	}
}

// StartElection 开始新的选举。
func (node *Node) StartElection() {
	node.State.ElectionResults = make(map[int]bool)
	// 增加提案编号并开始选举
	node.State.ProposalNumber++
	node.State.ProposalValue = node.ID
	electionMessage := Message{
		Type:   MessageTypePrepare,
		Number: node.State.ProposalNumber,
		Value:  node.State.ProposalValue, // 如果没有值，可以设置为nil
		From:   node.ID,
	}
	// 发送选举消息给所有其他节点
	node.Env.BroadcastMessage(node, electionMessage)
	fmt.Println("Node", node.ID, "started election with proposal number", node.State.ProposalNumber)
}

// HandlePromise 处理来自Acceptor的Promise响应。
func (node *Node) HandlePromise(msg Message) {
	if msg.Number != node.State.ProposalNumber {
		return
	}

	// 更新选举结果
	node.State.ElectionResults[msg.From] = true
	// 如果获得多数承诺，开始Accept阶段
	if node.checkQuorum() {
		node.SendAccept()
	}
}

// SendAccept 发送Accept请求给Acceptor。
func (node *Node) SendAccept() {
	acceptMessage := Message{
		Type:   MessageTypeAccept,
		Number: node.State.ProposalNumber,
		Value:  node.State.ProposalValue,
		From:   node.ID,
	}
	// 发送Accept请求给所有Acceptor
	node.Env.BroadcastMessage(node, acceptMessage)
}

// HandleAccepted 处理来自Acceptor的Accepted响应。
func (node *Node) HandleAccepted(msg Message) {
	// 更新选举结果并检查是否获得多数节点的接受
	node.State.ElectionResults[msg.From] = true

	// 如果获得多数接受，提案被接受
	if node.checkQuorum() {
		node.ProposalAccepted(msg)
	}
}

// ProposalAccepted 处理提案被接受的情况。
func (node *Node) ProposalAccepted(msg Message) {
	// 广播提案被接受的消息
	chosenMessage := Message{
		Type:   MessageTypeChosen,
		Number: msg.Number,
		Value:  msg.Value,
		From:   node.ID,
	}
	node.Env.BroadcastMessage(node, chosenMessage)
}

// checkQuorum 检查是否获得多数节点的承诺或接受。
func (node *Node) checkQuorum() bool {
	// 这里应该有逻辑来检查ElectionResults中是否有足够的承诺或接受
	// 如果是，返回true，否则返回false
	len := len(node.Env.Nodes)
	enough := (len + 1) / 2

	count := 0
	for i := 0; i < len; i++ {
		if node.State.ElectionResults[i] == true {
			count++
		}
	}
	return count >= enough
}
