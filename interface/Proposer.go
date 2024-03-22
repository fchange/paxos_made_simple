package _interface

import "paxos_made_simple/model"

// Proposer 接口定义了提议者需要实现的方法。
type Proposer interface {
	StartElection() // 开始新的选举
	// HandlePromise 处理来自Acceptor的Promise响应。
	HandlePromise(msg model.Message)

	// SendAccept 发送Accept请求给Acceptor。
	SendAccept()

	// HandleAccepted 处理来自Acceptor的Accepted响应。
	HandleAccepted(msg model.Message)
}
