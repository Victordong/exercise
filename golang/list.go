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
	head := new(ListNode)
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var result *ListNode
	var cur, last, nextLast *ListNode = nil, nil, head
	pre := head
	total := k
	for pre != nil {
		cur = nil
		for pre != nil && total > 0 {
			temp := pre.Next
			pre.Next = cur
			cur = pre
			pre = temp
			total -= 1
		}
		if total == 0 && last != nil {
			last.Next = cur
			last = nextLast
			nextLast = pre
		} else if total == 0 {
			last = nextLast
			result = cur
			nextLast = pre
		} else {
			pre, cur = cur, nil
			for pre != nil {
				temp := pre.Next
				pre.Next = cur
				cur = pre
				pre = temp
			}
			if last != nil {
				last.Next = cur
			} else {
				result = cur
			}
		}
		total = k
	}
	return result
}

func partMerge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	iter := 1
	for iter < n {
		for i := 0; i+iter < n; i += iter * 2 {
			lists[i] = partMerge(lists[i+iter], lists[i])
		}
		iter *= 2
	}
	return lists[0]
}
