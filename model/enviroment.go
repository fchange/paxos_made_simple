package model

import (
	"math/rand"
)

// Environment 表示Paxos算法的运行环境。
type Environment struct {
	Nodes map[int]*Node // 节点切片
}

// NewEnvironment 创建并初始化一个新的Environment实例。
func NewEnvironment(nodeCount int) *Environment {
	// 初始化节点和通信channel
	env := &Environment{
		Nodes: make(map[int]*Node, nodeCount),
	}
	for i := 0; i < nodeCount; i++ {
		env.Nodes[i] = NewNode(i, env)
	}
	return env
}

// RunPaxos 启动Paxos算法的执行。
func (env *Environment) RunPaxos() {
	// 启动所有节点的goroutine
	for _, node := range env.Nodes {
		go func(node *Node) {
			node.ParticipateInPaxos()
		}(node)
	}
}

// ReceiveMessage 模拟节点A向节点B发送消息。
func (env *Environment) ReceiveMessage(nodeId int, msg Message) {
	node := env.Nodes[nodeId]
	node.Channel <- msg
}

// BroadcastMessage 模拟节点A向所有其他节点广播消息。
func (env *Environment) BroadcastMessageDetail(node *Node, messageType MessageType, number int, value interface{}) {
	message := NewMessage(messageType, number, value, node.ID)
	for _, n := range env.Nodes {
		if n.ID != node.ID {
			n.Channel <- *message
		}
	}
}

func (env *Environment) BroadcastMessage(node *Node, message Message) {
	for _, n := range env.Nodes {
		if n.ID != node.ID {
			n.Channel <- message
		}
	}
}

// ReceiveClientMessage 模拟客户端向环境发送消息，并随机选择一个节点来接收。
func (env *Environment) ReceiveClientMessage(msg string) {
	message := NewLogMessage(msg)
	// 随机选择一个节点来接收客户端消息
	randNodeIndex := rand.Intn(len(env.Nodes))
	randNode := env.Nodes[randNodeIndex]
	// 将客户端消息发送到随机选择的节点
	randNode.Channel <- *message
}
