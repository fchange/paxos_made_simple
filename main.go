package main

import (
	"paxos_made_simple/model"
)

func main() {
	env := model.NewEnvironment(3) // 创建一个包含3个节点的环境
	env.RunPaxos()                 // 运行Paxos算法

	env.ReceiveClientMessage(*model.NewLogMessage("Hello, world!"))
}
