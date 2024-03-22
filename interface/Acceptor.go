package _interface

import "paxos_made_simple/model"

// Acceptor 接口定义了接受者需要实现的方法。
type Acceptor interface {
	// HandlePrepare 处理来自Proposer的Prepare请求。
	HandlePrepare(msg model.Message)

	// HandleAccept 处理来自Proposer的Accept请求。
	HandleAccept(msg model.Message)
}
