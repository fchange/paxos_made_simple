package model

import "fmt"

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

func (node *Node) ParticipateInPaxos(commChannel chan Message) {
	for {
		select {
		case msg := <-node.Channel:
			// 处理来自其他节点的消息
			fmt.Println("Node", node.ID, "received message:", msg)
		case msg := <-commChannel:
			// 处理来自客户端的请求
			fmt.Println("Node", node.ID, "received message:", msg)
		}
	}

	node.Background()
}

// 在分布式系统中，"看门狗"（Watchdog）机制是一种常见的方法，用于检测节点（尤其是领导者）是否已经失败或无法正常工作，并在必要时触发选举来选择新的领导者。
// 每个节点都会启动一个看门狗定时器，当定时器超时时，节点会认为领导者已经失败，并开始新的选举。
func (Node *Node) Background() {
	// TODO 看门狗机制
}
