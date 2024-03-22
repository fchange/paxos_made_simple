package model

import "fmt"

// Message 是节点间通信的消息结构体。
type Message struct {
	Type   MessageType
	Number int
	Value  interface{}
	From   int
}

// MessageType 定义了消息的类型。
type MessageType int

// 定义消息类型的常量
const (
	MessageTypePrepare    MessageType = iota // 准备阶段消息
	MessageTypePromise                       // 承诺阶段消息
	MessageTypeAccept                        // 接受阶段消息
	MessageTypeAccepted                      // 接受成功消息
	MessageTypeChosen                        // 选中值消息
	MessageTypeProposal                      // 提案消息
	MessageTypeWithdrawal                    // 撤回提案消息
	MessageTypeHeartBeat
	MessageTypeMonitorLeader
)

// String 方法为MessageType提供字符串表示。
func (mt MessageType) String() string {
	switch mt {
	case MessageTypePrepare:
		return "Prepare"
	case MessageTypePromise:
		return "Promise"
	case MessageTypeAccept:
		return "Accept"
	case MessageTypeAccepted:
		return "Accepted"
	case MessageTypeChosen:
		return "Chosen"
	case MessageTypeProposal:
		return "Proposal"
	case MessageTypeWithdrawal:
		return "Withdrawal"
	case MessageTypeHeartBeat:
		return "HeartBeat"
	case MessageTypeMonitorLeader:
		return "MonitorLeader"
	default:
		return fmt.Sprintf("Unknown message type: %d", mt)
	}
}

// NewMessage 创建并返回一个新的Message实例。
func NewMessage(t MessageType, n int, v interface{}, from int) *Message {
	return &Message{
		Type:   t,
		Number: n,
		Value:  v,
		From:   from,
	}
}

// NewPromiseMessage 创建并返回一个新的TypePromise消息实例。
func NewPromiseMessage(number int) *Message {
	return NewMessage(MessageTypePromise, number, nil, -1)
}

// NewLogMessage 创建并返回一个新的TypePrepare消息实例，用于客户端提案。
// 客户端不需要指定提案编号，节点将在处理消息时分配编号。
func NewLogMessage(value interface{}) *Message {
	// 假设提案编号（n）由节点在处理消息时分配，客户端不需要提供。
	return NewMessage(MessageTypePrepare, 0, value, -1) // from设置为-1表示这是一个来自客户端的消息
}

// String 方法返回Message的格式化字符串表示。
func (m *Message) String() string {
	return fmt.Sprintf("Type: %s, Number: %d, Value: %v, From: %d", m.Type.String(), m.Number, m.Value, m.From)
}
