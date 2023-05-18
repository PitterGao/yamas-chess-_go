# yamas-chess-_go

# 亚马逊棋（Amazon Chess）AI

本项目实现了基于蒙特卡洛树搜索算法的亚马逊棋AI。亚马逊棋是一种棋类游戏，既有围棋、象棋的某些规则，又有独特的规则和特色。

## 算法原理

蒙特卡洛树搜索算法是一种搜索优化技术，它通过随机模拟来评估每个决策的价值，并逐步构建一棵搜索树。在搜索树上不断扩展和更新节点，从而找到最优解。

该算法主要由以下四个阶段组成：

1. 选择（Selection）：从根节点出发，依据一定规则选择一个子节点。
2. 模拟（Simulation）：对选定的子节点进行一次模拟。
3. 扩展（Expansion）：根据模拟结果，在搜索树上扩展新的节点。
4. 回溯（Backpropagation）：将模拟得到的信息从叶子节点传递回根节点，并更新搜索树。

## 项目结构

```
amazon-chess-ai/
├── amazon/
│   ├── board.go        # 游戏棋盘定义
│   ├── moves.go        # 走棋算法实现
│   └── rules.go        # 游戏规则定义
├── mcts/
│   ├── mcts.go         # MCTS算法实现
│   └── node.go         # 搜索树节点定义
├── main.go             # 主程序入口
└── README.md           # 项目说明文档
```

## 使用方法

1. 克隆本仓库到本地。
2. 在终端进入项目根目录，运行`go run main.go`命令启动程序。
3. 根据提示输入游戏信息（如棋盘大小、胜利条件等）。
4. 开始游戏并与AI对弈。

## 参考资料

- Müller, M., & Schaeffer, J. (2002). The game of Amazons: a survey. Games of No Chance 3, 157-180.
- Browne, C. B., Powley, E., Whitehouse, D., Lucas, S. M., Cowling, P. I., & Rohlfshagen, P. (2012). A survey of Monte Carlo tree search methods. IEEE Transactions on Computational Intelligence and AI in games, 4(1), 1-43.

## 版权信息

本项目基于 MIT 协议开源。
