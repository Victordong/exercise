package main

import "fmt"

func main() {
	root := new(ListNode)
	root.Val = 1
	cur := root
	for i := 2; i <= 5; i++ {
		newNode := new(ListNode)
		newNode.Val = i
		cur.Next = newNode
		cur = newNode
	}
	cur = reverseKGroup(root, 3)
	fmt.Println(cur.Val)
}
