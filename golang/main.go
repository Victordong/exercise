package main

import "fmt"

func main() {
	A := []int{3, 9, 20, 15, 7}
	B := []int{9, 3, 15, 20, 7}
	root := buildTree(A, B)
	fmt.Println(levelOrder(root))
}
