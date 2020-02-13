package main

import (
	"fmt"
	"sort"
)

func position(nums []int, left int, right int) int {
	temp := nums[left]
	for left < right {
		for left < right && nums[right] <= temp {
			right -= 1
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= temp {
			left += 1
		}
		nums[right] = nums[left]
	}
	nums[left] = temp
	return left
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pos := position(nums, left, right)
	fmt.Println(left, right, pos)
	quickSort(nums, left, pos-1)
	quickSort(nums, pos+1, right)
}

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			break
		}
		pos := position(nums, left, right)
		if pos < k-1 {
			left = pos + 1
		} else if pos > k-1 {
			right = pos - 1
		} else {
			return nums[pos]
		}
	}
	return nums[k-1]
}

type meeting struct {
	begin int
	end   int
}

type meetingArr []meeting

func (m meetingArr) Len() int {
	return len(m)
}

func (m meetingArr) Less(i, j int) bool {
	return m[i].begin <= m[j].begin
}

func (m meetingArr) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func canAttendMeetings(intervals [][]int) bool {
	m := len(intervals)
	result := make(meetingArr, 0)
	for i := 0; i < m; i++ {
		result = append(result, meeting{begin: intervals[i][0], end: intervals[i][1]})
	}
	sort.Sort(result)
	for i := 1; i < m; i++ {
		if result[i].begin < result[i-1].begin {
			return false
		}
	}
	return true
}

func minMeetingRooms(intervals [][]int) int {
	m := len(intervals)
	if m == 0 {
		return 0
	}
	result := make(meetingArr, 0)
	for i := 0; i < m; i++ {
		result = append(result, meeting{begin: intervals[i][0], end: intervals[i][1]})
	}
	sort.Sort(result)
	minRooms := 0
	endStack := make([]int, 0)
	helpStack := make([]int, 0)
	endStack = append(endStack, result[0].end)
	for i := 1; i < m; i++ {
		for len(endStack) != 0 && result[i].begin >= endStack[len(endStack)-1] {
			endStack = endStack[:len(endStack)-1]
		}
		for len(endStack) != 0 && result[i].end > endStack[len(endStack)-1] {
			helpStack = append(helpStack, endStack[len(endStack)-1])
			endStack = endStack[:len(endStack)-1]
		}
		endStack = append(endStack, result[i].end)
		for len(helpStack) != 0 {
			endStack = append(endStack, helpStack[len(helpStack)-1])
			helpStack = helpStack[:len(helpStack)-1]
		}
	}
	minRooms = max(minRooms, len(endStack))
	return minRooms
}
