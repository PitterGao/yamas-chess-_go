module github.com/PitterGao/MCTS

go 1.19

require github.com/PitterGao/Regulation v0.0.0-20220723102016-321ef0335ad3

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)
replace (
	github.com/PitterGao/Regulation => ../Regulation
)