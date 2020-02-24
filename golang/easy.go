package golang

import (
	"fmt"
	"sort"
	"strconv"
)

func kthSmallest1(matrix [][]int, k int) int {
	n := len(matrix)
	var i int
	var j int
	if k%n == 0 {
		i = k/n - 1
		j = n
	} else {
		i = k/n - 1
		j = k - i*n - 1
	}
	return matrix[i][j]
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	hashMap := make(map[int]int)
	for i := 0; i < n; i++ {
		hashMap[nums[i]] = i
	}
	for i := 0; i < n; i++ {
		if _, ok := hashMap[target-nums[i]]; ok {
			if i != hashMap[target-nums[i]] {
				return []int{i, hashMap[target-nums[i]]}
			}
		}
	}
	return nil
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	total := 0
	for {
		total = 10*total + temp%10
		temp = temp / 10
		if temp == 0 {
			break
		}
	}
	if total == x {
		return true
	}
	return false
}

func reverse1(x int) int {
	total := 0
	re := false
	if x < 0 {
		re = true
		x = x * -1
	}
	for {
		total = total*10 + x%10
		x = x / 10
		if x == 0 {
			break
		}
	}
	if re {
		total = total * -1
	}
	return total
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i += 1
		}
		if matrix[i][j] > target {
			j -= 1
		}
	}
	return false
}

type largestNumberList []int

func (l largestNumberList) Len() int {
	return len(l)
}

func (l largestNumberList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l largestNumberList) Less(i, j int) bool {
	return strconv.Itoa(l[i])+strconv.Itoa(l[j]) > strconv.Itoa(l[j])+strconv.Itoa(l[i])
}

func largestNumber(nums []int) string {
	l := make(largestNumberList, len(nums))
	copy(l, nums)
	sort.Sort(l)
	result := ""
	for _, value := range l {
		result += strconv.Itoa(value)
	}
	part, err := strconv.Atoi(result)
	if err != nil {
		return result
	} else {
		return strconv.Itoa(part)
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	i, j, cur := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[i] {
			nums1[cur] = nums1[i]
			i -= 1
		} else {
			nums1[cur] = nums2[j]
			j -= 1
		}
		cur -= 1
	}
	if j > 0 {
		for i := 0; i <= j; i++ {
			nums1[i] = nums2[i]
		}
	}
	fmt.Println(nums1)
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func rotate(nums []int, k int) {
	n := len(nums)
	total := 0
	var temp, cur int
	for start := 0; total < n; start++ {
		cur, temp = start, nums[start]
		for {
			nTemp := nums[(cur+k)%n]
			nums[(cur+k)%n] = temp
			temp = nTemp
			cur = (cur + k) % n
			total = total + 1
			if cur == start {
				break
			}
		}
	}
}
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int][]int)
	for index, num := range nums {
		if value, ok := hashMap[num]; ok {
			for i := len(value) - 1; i >= 0; i-- {
				if (index - value[i]) <= k {
					return true
				} else {
					hashMap[num] = append(hashMap[num], index)
				}
			}
		} else {
			hashMap[num] = []int{index}
		}
	}
	return false
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	now := 0
	i, j := 0, 0
	result := make([]int, 0)
	for {
		if m-now*2 <= 1 || n-now*2 <= 1 {
			break
		}
		for ; j < n-1-now; j++ {
			result = append(result, matrix[i][j])
		}
		for ; i < m-1-now; i++ {
			result = append(result, matrix[i][j])
		}
		for ; j > now; j-- {
			result = append(result, matrix[i][j])
		}
		for ; i > now; i-- {
			result = append(result, matrix[i][j])
		}
		now += 1
		i, j = i+1, j+1
	}
	fmt.Println(n - now)
	if m-now*2 == 1 {
		for ; j < n-now; j++ {
			result = append(result, matrix[i][j])
		}
	} else if n-now*2 == 1 {
		for ; i < n-now; i++ {
			result = append(result, matrix[i][j])
		}
	}
	return result
}

type MinStack struct {
	main []int
	help []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{main: make([]int, 0), help: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.main) == 0 {
		this.help = append(this.help, x)
	} else {
		var top = this.main[len(this.main)-1]
		if x >= top {
			this.help = append(this.help, x)
		}
	}
	this.main = append(this.main, x)
}

func (this *MinStack) Pop() {
	var top = this.main[len(this.main)-1]
	if top == this.help[len(this.main)-1] {
		this.help = this.help[:len(this.help)-1]
	}
	this.main = this.main[:len(this.main)-1]
}

func (this *MinStack) Top() int {
	return this.main[len(this.main)-1]
}

func (this *MinStack) GetMin() int {
	return this.help[len(this.help)-1]
}

func checkValidString(s string) bool {
	mainStack := make([]int, 0)
	starStack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			mainStack = append(mainStack, i)
		} else if s[i] == ')' {
			if len(mainStack) > 0 {
				mainStack = mainStack[:len(mainStack)-1]
			} else {
				if len(starStack) > 0 {
					starStack = starStack[:len(starStack)-1]
				} else {
					return false
				}
			}
		} else {
			starStack = append(starStack, i)
		}
	}
	for len(mainStack) != 0 {
		if len(starStack) == 0 {
			return false
		}
		if starStack[len(starStack)-1] < mainStack[len(starStack)-1] {
			starStack = starStack[:len(starStack)-1]
			mainStack = mainStack[:len(mainStack)-1]
		} else {
			return false
		}
	}
	return true
}

func checkValidString1(s string) bool {
	minV, maxV := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			minV, maxV = minV+1, maxV+1
		} else if s[i] == ')' {
			if maxV <= 0 {
				return false
			}
			minV, maxV = minV-1, maxV-1
		} else {
			minV = minV - 1
			maxV = maxV + 1
		}
	}
	if minV <= 0 && maxV >= 0 {
		return true
	}
	return false
}

func isUgly(num int) bool {
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%5 == 0 {
			num /= 5
		} else if num%3 == 0 {
			num /= 3
		} else {
			return false
		}
	}
	return true
}
