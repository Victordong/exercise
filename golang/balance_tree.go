package golang

type BalanceNode struct {
	Childrens []*BalanceNode
}

type BalanceTree struct {
	root *BalanceNode
}

func(tree *BalanceTree) insertNode(value int) {
}
