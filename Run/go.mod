module Winrate

go 1.19

require (
	github.com/PitterGao/MCTS v0.0.0-20220723114651-80b546fbf692 // indirect
	github.com/PitterGao/Regulation v0.0.0-20220723102016-321ef0335ad3 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)
replace (
	github.com/PitterGao/MCTS => ../Mcts
)

replace (
	github.com/PitterGao/Regulation => ../Regulation
)