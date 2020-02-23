package golang

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

func findKthLargest1(nums []int, k int) int {
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

type mergeInterval struct {
	begin int
	end   int
}

type mergeIntervalList []mergeInterval

func (m mergeIntervalList) Len() int {
	return len(m)
}

func (m mergeIntervalList) Less(i, j int) bool {
	return m[i].begin <= m[j].begin
}

func (m mergeIntervalList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	list := make(mergeIntervalList, 0)
	for i := 0; i < n; i++ {
		list = append(list, mergeInterval{begin: intervals[i][0], end: intervals[i][1]})
	}
	sort.Sort(list)
	result := make([][]int, 0)
	for i := 1; i < n; i++ {
		if list[i].begin <= list[i-1].end {
			list[i].begin = list[i-1].begin
			list[i].end = max(list[i-1].end, list[i].end)
		} else {
			result = append(result, []int{list[i-1].begin, list[i-1].end})
		}
	}
	result = append(result, []int{list[len(list)-1].begin, list[len(list)-1].end})
	return result
}

func partHeapSort(arr []int, begin int, end int) {
	father := begin
	son := father*2 + 1
	for son <= end {
		if son+1 <= end && arr[son] > arr[son+1] {
			son = son + 1
		}
		if arr[father] < arr[son] {
			return
		} else {
			arr[father], arr[son] = arr[son], arr[father]
			father = son
			son = father * 2
		}
	}
}

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		partHeapSort(arr, i, n-1)
	}
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		partHeapSort(arr, 0, i-1)
	}
	fmt.Println(arr)
}

//func countingSort(arr []int) {
//	minNum, maxNum := arr[0], arr[0]
//	for
//}

func quickFind(nums []int, left int, right int) int {
	i, j, temp := left, right, nums[left]
	for i < j {
		for nums[j] >= temp && i < j {
			j--
		}
		nums[i] = nums[j]
		for nums[i] <= temp && i < j {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = temp
	return i
}

func findKthLargest2(nums []int, k int) int {
	left, right, pos := 0, len(nums)-1, 0
	for {
		pos = quickFind(nums, left, right)
		if pos == k {
			return nums[k]
		} else if pos > k {
			right = pos
		} else {
			left = pos
		}
	}
}

func heapFindKthLargest(nums []int, begin int, end int) {
	father := begin
	son := father*2 + 1
	for son <= end {
		if son+1 <= end && nums[son] < nums[son+1] {
			son = son + 1
		}
		if nums[father] > nums[son] {
			return
		} else {
			nums[father], nums[son] = nums[son], nums[father]
			father = son
			son = father*2 + 1
		}
	}
}

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapFindKthLargest(nums, i, len(nums)-1)
	}
	total := 1
	for total < k {
		nums[0], nums[len(nums)-total] = nums[len(nums)-total], nums[0]
		heapFindKthLargest(nums, 0, len(nums)-total-1)
	}
	return nums[0]
}

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := left; i <= right; {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left, i = left+1, i+1
		} else if nums[i] == 1 {
			i = i + 1
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right = right - 1
		}
	}
}
