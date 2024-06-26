# 分布式系统中的 Paxos 算法实现

这是一个简化的 Paxos 算法的 Go 语言实现，用于教学和学习目的。Paxos 算法是一种解决分布式系统中的共识问题的算法，能够在部分节点失败的情况下，保证系统中的多个节点达成一致的决策。
参考自：https://lamport.azurewebsites.net/pubs/paxos-simple.pdf

## 功能

- 实现了基本的 Paxos 算法，包括提案、承诺、接受和选定值的通信过程。
- 引入了随机化选举超时，以减少多个节点同时发起选举的可能性。
- 集成了看门狗机制，用于自动检测领导者节点的状态，并在领导者离线时触发新的选举。

## 快速开始

1. 确保你的系统上安装了 Go 语言环境。
2. 克隆本项目到本地：
   ```
   git clone https://github.com/fchange/paxos_made_simple
   ```
3. 进入项目目录并安装依赖：
   ```
   cd paxos_made_simple
   go mod tidy
   ```
4. 运行项目：
   ```
   go run main.go
   ```

## 架构

本项目包含以下主要组件：

- `Node`：表示参与 Paxos 算法的节点。每个节点都有自己的状态和通信通道。
- `Environment`：表示整个分布式环境，包含所有节点和用于节点间通信的通道。
- `Message`：表示节点间交换的消息，包括提案、承诺、接受等。

## 贡献

欢迎对本项目做出贡献。如果你发现了问题或有改进的建议，请通过以下方式联系我们：

- 提交 [issue](https://github.com/fchange/paxos_made_simple/issues) 来报告问题。
- 发送 [pull request](https://github.com/fchange/paxos_made_simple/pulls) 来贡献代码。

## 许可证

本项目使用 MIT 许可证。详情请参阅 `LICENSE` 文件。

---

请根据你的项目具体情况调整上述内容。如果你的项目有特定的安装步骤、配置选项或其他重要信息，确保在 `README.md` 文件中包含这些信息。