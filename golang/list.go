package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addNum := 0
	head := l1
	cur1, cur2 := l1, l2
	pre := cur1
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val + addNum
		addNum = 0
		if sum >= 10 {
			addNum = 1
			sum = sum - 10
		}
		cur1.Val = sum
		cur2.Val = sum
		pre = cur1
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur2 != nil {
		head = l2
		cur1 = cur2
	}
	for cur1 != nil {
		sum := cur1.Val + addNum
		addNum = 0
		if sum >= 10 {
			addNum = 1
			sum = sum - 10
		}
		cur1.Val = sum
		pre = cur1
		cur1 = cur1.Next
	}
	if addNum != 0 {
		newNode := new(ListNode)
		newNode.Val = addNum
		pre.Next = newNode
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode = nil
	pre := head
	for pre != nil {
		temp := pre.Next
		pre.Next = cur
		cur = pre
		pre = temp
	}
	return cur
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur1, cur2, pre, temp *ListNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		head, cur1, cur2, pre = l1, l1, l2, l1
	} else {
		head, cur1, cur2, pre = l2, l2, l1, l2
	}
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			pre, cur1 = cur1, cur1.Next
		} else {
			pre.Next, temp, pre = cur2, cur2.Next, cur2
			cur2.Next, cur2 = cur1, temp
		}
	}
	if cur1 == nil {
		pre.Next = cur2
	} else if cur2 == nil {
		pre.Next = cur1
	}
	return head
}
