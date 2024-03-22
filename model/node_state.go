package model

import "time"

// NodeState 表示节点的状态，增加了与选举相关的字段。
type NodeState struct {
	ProposalNumber  int          // 当前节点提议的编号
	ProposalValue   interface{}  // 当前节点提议的值
	Leader          int          // 当前领导者的ID，如果节点是领导者则为自身ID
	ElectionResults map[int]bool // 选举结果，记录其他节点是否投票给了当前节点
	// 其他状态相关的字段...

	LastHeartBeatTime time.Time // 记录最后一次接收到领导者消息的时间
}
