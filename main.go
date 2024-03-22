package main

import (
	"fmt"
	"paxos_made_simple/model"
)

func main() {
	env := model.NewEnvironment(3) // 创建一个包含3个节点的环境
	env.RunPaxos()                 // 运行Paxos算法

	// 读取控制台输入
	var input string
	fmt.Scanln(&input)
	env.ReceiveClientMessage(input)
}
