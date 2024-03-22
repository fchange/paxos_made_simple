package model

import (
	"math/rand"
	"sync"
)

// Environment 表示Paxos算法的运行环境。
type Environment struct {
	Nodes       []*Node         // 节点切片
	CommChannel chan Message    // 用于节点间通信的channel
	wg          *sync.WaitGroup // 用于等待所有goroutine完成
}

// NewEnvironment 创建并初始化一个新的Environment实例。
func NewEnvironment(nodeCount int) *Environment {
	// 初始化节点和通信channel
	env := &Environment{
		Nodes:       make([]*Node, nodeCount),
		CommChannel: make(chan Message, 100), // 假设bufferSize为100
		wg:          &sync.WaitGroup{},
	}
	for i := 0; i < nodeCount; i++ {
		env.Nodes[i] = &Node{
			ID:      i,
			Env:     env,
			Channel: make(chan Message),
		}
	}
	return env
}

// RunPaxos 启动Paxos算法的执行。
func (env *Environment) RunPaxos() {
	// 启动所有节点的goroutine
	for _, node := range env.Nodes {
		env.wg.Add(1)
		go func(node *Node) {
			defer env.wg.Done()
			node.ParticipateInPaxos(env.CommChannel)
		}(node)
	}
	// 发送初始的客户端请求到Paxos环境
	env.CommChannel <- Message{Type: TypeClientRequest, From: -1, Value: "Initial client request"}
	// 等待所有节点完成
	env.wg.Wait()
}

// ReceiveMessage 模拟节点A向节点B发送消息。
func (env *Environment) ReceiveMessage(toNode *Node, msg Message) {
	toNode.Channel <- msg // 将消息发送到指定节点的通道
}

// BroadcastMessage 模拟节点A向所有其他节点广播消息。
func (env *Environment) BroadcastMessage(fromNode *Node, msg Message) {
	for _, node := range env.Nodes {
		if node != fromNode { // 避免向发送者自身发送消息
			node.Channel <- msg
		}
	}
}

// ReceiveClientMessage 模拟客户端向环境发送消息，并随机选择一个节点来接收。
func (env *Environment) ReceiveClientMessage(msg Message) {
	// 随机选择一个节点来接收客户端消息
	randNodeIndex := rand.Intn(len(env.Nodes))
	randNode := env.Nodes[randNodeIndex]
	// 将客户端消息发送到随机选择的节点
	randNode.Channel <- msg
}
