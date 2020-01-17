package main

import "fmt"

func main() {
	b := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   1,
		Left:  b,
		Right: nil,
	}

	fmt.Println(countNodes(a))
}
