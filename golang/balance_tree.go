package golang

import "reflect"

type BalanceNode struct {
	Childrens []*BalanceNode
	IsLeaf    bool
	Value     interface{}
	Kind      reflect.Kind
}

type BalanceTree struct {
	root *BalanceNode
}

func (tree *BalanceTree) InsertNode(value interface{}) {
}

func (tree *BalanceTree) DeleteNodesWithValue(value interface{}) {

}

func (tree *BalanceTree) InitBalanceTree() {
	tree.root = new(BalanceNode)
}

func (tree *BalanceTree) Clear() {}
