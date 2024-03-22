package model

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
	TypePrepare MessageType = iota
	TypePromise
	TypeAccept
	TypeAccepted
	TypeChosen
	TypeClientRequest
)

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
	return NewMessage(TypePromise, number, nil, -1)
}

// NewLogMessage 创建并返回一个新的TypePrepare消息实例，用于客户端提案。
// 客户端不需要指定提案编号，节点将在处理消息时分配编号。
func NewLogMessage(value interface{}) *Message {
	// 假设提案编号（n）由节点在处理消息时分配，客户端不需要提供。
	return NewMessage(TypePrepare, 0, value, -1) // from设置为-1表示这是一个来自客户端的消息
}
