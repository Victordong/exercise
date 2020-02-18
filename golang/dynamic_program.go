package golang

import "math"

func longestPalindrome(s string) string {
	var a [1001][1001]int
	n := len(s)
	if n == 0 {
		return s
	}
	for i := 0; i < n; i++ {
		a[i][i] = 1
	}
	begin, end := 0, 0
	maxLen := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			a[i-1][i] = 2
		} else {
			a[i-1][i] = 0
		}
		if maxLen < a[i-1][i] {
			begin = i - 1
			end = i
			maxLen = a[i-1][i]
		}
	}
	for m := 2; m < n; m++ {
		for i := 0; i < n; i++ {
			j := m + i
			if j >= n {
				break
			}
			if s[i] == s[j] && a[i+1][j-1] != 0 {
				a[i][j] = a[i+1][j-1] + 2
				if maxLen < a[i][j] {
					begin = i
					end = j
					maxLen = a[i][j]
				}
			} else {
				a[i][j] = 0
			}
		}
	}
	return s[begin : end+1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var a [101][101]int
	can := true
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			can = false
		}
		if can {
			a[i][0] = 1
		} else {
			a[i][0] = 0
		}
	}
	can = true
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			can = false
		}
		if can {
			a[0][j] = 1
		} else {
			a[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				a[i][j] = 0
			} else {
				a[i][j] = a[i-1][j] + a[i][j-1]
			}
		}
	}
	return a[m-1][n-1]
}

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][len(triangle[i])-1] += triangle[i-1][len(triangle[i])-2]
		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	min := int(^uint(0) >> 1)
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		if min > triangle[len(triangle)-1][j] {
			min = triangle[len(triangle)-1][j]
		}
	}
	return min
}

func min(first int, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func max(first int, second int) int {
	if first < second {
		return second
	} else {
		return first
	}
}

func maxProduct1(nums []int) int {
	var imax = 1
	var imin = 1
	var maxNum = ^(int(^uint(0) >> 1))
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := imax
			imax = imin
			imin = temp
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		maxNum = max(maxNum, imax)
	}
	return imax
}

func maxProfit1(prices []int) int {
	min := prices[0]
	max := ^min
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	if max < 0 {
		return 0
	} else {
		return max
	}
}

func maxProfit2(prices []int) int {
	n := len(prices)
	var b [100000]int
	b[0] = 0
	for i := 1; i < n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			var val int
			if prices[i] > prices[j] {
				val = 0
			} else {
				val = prices[j] - prices[i]
			}
			if max < b[j]+val {
				max = b[j] + val
			}
		}
		b[i] = max
	}
	return b[n-1]
}

func maxProfit3(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	max := 0
	for i := 0; i < n-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}

func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n < 4 {
		maxNum := 0
		for i := 0; i < n; i++ {
			if maxNum < nums[i] {
				maxNum = nums[i]
			}
		}
		return maxNum
	}
	maxWithFirst := nums[2]
	maxWithoutFirst := nums[1]
	pre := 0
	for i := 2; i < n; i++ {
		temp := maxWithoutFirst
		maxWithoutFirst = max(maxWithoutFirst, pre+nums[i])
		pre = temp
	}
	pre = 0
	for i := 3; i < n; i++ {
		temp := maxWithFirst
		maxWithFirst = max(maxWithFirst, pre+nums[i])
		pre = temp
	}
	return max(maxWithFirst+nums[0], maxWithoutFirst)
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	var dp [1000][1000]int
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}
	maxLength := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			length := min(dp[i-1][j], dp[i][j-1])
			if matrix[i-length][j-length] == '1' && matrix[i][j] == '1' {
				dp[i][j] = length + 1
				if maxLength < dp[i][j] {
					maxLength = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return maxLength * maxLength
}

func maxProfit4(prices []int) int {
	n := len(prices)
	var curBuy int
	var curSell int
	var preSell int
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	preSell = 0
	if prices[1] > prices[0] {
		curSell = prices[1] - prices[0]
		curBuy = -1 * prices[0]
	} else {
		curSell = 0
		curBuy = -1 * prices[1]
	}
	for i := 2; i < n; i++ {
		tempSell := curSell
		curSell = max(curBuy+prices[i], curSell)
		curBuy = max(curBuy, preSell-prices[i])
		preSell = tempSell
	}
	return curSell
}

func maxProfit5(prices []int, fee int) int {
	n := len(prices)
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	var curBuy int
	var curSell int
	if prices[0]+fee < prices[1] {
		curSell = prices[1] - prices[0] - fee
		curBuy = -1*prices[0] - fee
	} else if prices[0] < prices[1] {
		curSell = 0
		curBuy = -1*prices[0] - fee
	} else {
		curSell = 0
		curBuy = -1*prices[1] - fee
	}
	for i := 2; i < n; i++ {
		tempSell := curSell
		tempBuy := curBuy
		curBuy = max(tempSell-prices[i]-fee, tempBuy)
		curSell = max(tempBuy+prices[i], tempSell)
	}
	return curSell
}

func maxProfit6(prices []int) int {
	n := len(prices)
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	var curBuy [3]int
	var curSell [3]int
	curSell[0] = 0
	curBuy[0] = 0
	if prices[1] > prices[0] {
		for j := 1; j < 3; j++ {
			curSell[j] = prices[1] - prices[0]
			curBuy[j] = -1 * prices[0]
		}
	} else {
		for j := 1; j < 3; j++ {
			curSell[j] = 0
			curBuy[j] = -1 * prices[1]
		}
	}
	for i := 2; i < n; i++ {
		for j := 1; j < 3; j++ {
			curSell[j] = max(curBuy[j]+prices[i], curSell[j])
			curBuy[j] = max(curSell[j-1]-prices[i], curBuy[j])
		}
	}
	return curSell[2]
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	if k > n/2 {
		result := 0
		for i := 0; i < n-1; i++ {
			if prices[i] < prices[i+1] {
				result += prices[i+1] - prices[i]
			}
		}
		return result
	}
	var curBuy [1000]int
	var curSell [1000]int
	curSell[0] = 0
	curBuy[0] = 0
	if prices[1] > prices[0] {
		for j := 1; j <= k; j++ {
			curSell[j] = prices[1] - prices[0]
			curBuy[j] = -1 * prices[0]
		}
	} else {
		for j := 1; j <= k; j++ {
			curSell[j] = 0
			curBuy[j] = -1 * prices[1]
		}
	}
	for i := 2; i < n; i++ {
		for j := 1; j <= k; j++ {
			curSell[j] = max(curBuy[j]+prices[i], curSell[j])
			curBuy[j] = max(curSell[j-1]-prices[i], curBuy[j])
		}
	}
	return curSell[k]
}

func countSubstrings(s string) int {
	n := len(s)
	var dp [10][10]int
	count := 0
	for i := 0; i < n; i++ {
		dp[i][i] = 1
		count += 1
	}
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
			count += 1
		}
	}
	for m := 2; m < n; m++ {
		for i := 0; i < n; i++ {
			j := m + i
			if j >= n {
				break
			}
			if s[i] == s[j] && dp[i+1][j-1] == 1 {
				dp[i][j] = 1
				count += 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	return count
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	var dp [10000]int
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		minNum := amount + 1
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				minNum = min(minNum, dp[i-coin]+1)
			}
		}
		if minNum != amount+1 {
			dp[i] = minNum
		} else {
			dp[i] = -1
		}

	}
	return dp[amount]
}

func change(amount int, coins []int) int {
	var dp [5001]int
	n := len(coins)
	if amount == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}

	dp[0] = 1
	for i := 0; i <= amount; i++ {
		if i%coins[0] == 0 {
			dp[i] = 1
		} else {
			dp[i] = 0
		}
	}
	for j := 1; j < n; j++ {
		for i := 1; i <= amount; i++ {

			if i >= coins[j] {
				dp[i] = dp[i] + dp[i-coins[j]]
			}
		}
	}
	return dp[amount]
}

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	print(n)
	if n == 0 {
		return 0
	}
	var dp [1000]int
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		maxNum := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && maxNum < dp[j] {
				maxNum = dp[j]
			}
		}
		dp[i] = maxNum + 1
	}
	max := dp[0]
	for i := 1; i < n; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	if n == 0 {
		return 0
	}
	var dp [1000]int
	dp[0] = 0
	dp[1] = min(costs[0], min(costs[1], costs[2]))
	for i := 2; i <= n; i++ {
		var index1 int
		var index2 int
		if days[i-1] <= 7 {
			index1 = 0
			index2 = 0
		} else if days[i-1] <= 30 {
			j := i - 2
			index1 = i - 1
			for ; j >= 0; j-- {
				if days[i-1]-days[j] < 7 {
					index1 = j
				}
			}
			index2 = 0
		} else {
			j := i - 2
			index1 = i - 1
			for ; j >= 0; j-- {
				if days[i-1]-days[j] < 7 {
					index1 = j
				}
			}
			j = i - 2
			index2 = j
			for ; j >= 0; j-- {
				if days[i-1]-days[j] < 30 {
					index2 = j
				}
			}
		}
		dp[i] = min(dp[i-1]+costs[0], min(dp[index1]+costs[1], dp[index2]+costs[2]))
	}
	return dp[n]
}

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var dp [2500]int
	dp[0] = nums[0]
	maxLength := 1
	for m := 1; m < len(nums); m++ {
		var mid int
		i, j := 0, maxLength
		for i < j {
			mid = (i + j) / 2
			if dp[mid] < nums[m] {
				i = mid + 1
			} else {
				j = mid
			}
		}
		if j == maxLength && nums[m] != dp[maxLength-1] {
			dp[maxLength] = nums[m]
			maxLength += 1
		} else {
			dp[i] = nums[m]
		}
	}
	return maxLength
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	print(n)
	if n == 0 {
		return 0
	}
	var dp [1000]int
	var count [1000]int
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}
	for i := 1; i < n; i++ {
		maxNum := 0
		maxSum := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && maxNum < dp[j] {
				maxNum = dp[j]
				maxSum = count[j]
			} else if nums[i] > nums[j] && maxNum == dp[j] {
				maxSum += count[j]
			}
		}
		dp[i] = maxNum + 1
		if maxSum != 0 {
			count[i] = maxSum
		}
	}
	max := dp[0]
	number := count[1]
	for i := 1; i < n; i++ {
		if max < dp[i] {
			max = dp[i]
			number = count[i]
		} else if max == dp[i] {
			number += count[i]
		}
	}
	return number
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	var dp [1000][1000]int
	equal := false
	for i := 0; i < m; i++ {
		if text1[i] == text2[0] {
			equal = true
		}
		if !equal {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}
	equal = false
	for i := 0; i < n; i++ {
		if text2[i] == text1[0] {
			equal = true
		}
		if !equal {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func countSquares(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}
	var dp [301][301]int
	for i := 0; i < n; i++ {
		if matrix[i][0] == 1 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for i := 0; i < m; i++ {
		if matrix[0][i] == 1 {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			length := min(dp[i-1][j], dp[i][j-1])
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if matrix[i-length][j-length] == 1 {
					dp[i][j] = length + 1
				} else {
					dp[i][j] = length
				}
			}
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result += dp[i][j]
		}
	}
	return result
}

func minDistance1(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	var dp [1000][1000]int
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]))
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}
	return dp[n][m]
}

func climbStairs(n int) int {
	pre, cur := 1, 2
	if n == 0 {
		return 1
	} else if n == 1 {
		return pre
	} else if n == 2 {
		return cur
	}
	for i := 3; i <= n; i++ {
		temp := cur
		cur = cur + pre
		pre = temp
	}
	return cur
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n == 0 || n == 1 || n == 2 {
		return true
	}
	var dp [21][21]int
	dp[0][0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i][i] = nums[i]
		dp[i-1][i] = max(nums[i-1], nums[i]) - min(nums[i-1], nums[i])
	}
	for m := 2; m < n; m++ {
		for i := 0; i < n; i++ {
			j := i + m
			if j >= n {
				break
			}
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	if dp[0][n-1] >= 0 {
		return true
	}
	return false
}

func contains(s string, wordDict []string) bool {
	for _, str := range wordDict {
		if s == str {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	m := len(wordDict)
	if n == 0 {
		return true
	}
	if m == 0 {
		return false
	}
	var dp [1000]bool
	dp[0] = true
	for i := 1; i <= n; i++ {
		dp[i] = false
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			newString := s[i:j]
			if dp[i] == true && contains(newString, wordDict) {
				dp[j] = true
			}
		}
	}
	return dp[n]
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, 0)
	for i := 0; i < n; i++ {
		partDp := make([]int, n)
		dp = append(dp, partDp)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			dp[i-1][i] = 2
		} else {
			dp[i-1][i] = 1
		}
	}
	for i := n - 2; i >= 1; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	var total int
	if len(nums) < 2 {
		return len(nums)
	}
	var trend bool
	total = 1
	j := 1
	for ; j < n; j++ {
		if nums[j] < nums[j-1] {
			trend = false
			total = 2
			break
		} else if nums[j] > nums[j-1] {
			trend = true
			total = 2
			break
		}
	}
	for i := j + 1; i < n; i++ {
		if trend && nums[i] < nums[i-1] {
			total += 1
			trend = false
		}
		if !trend && nums[i] > nums[i-1] {
			total += 1
			trend = true
		}
	}
	return total
}

func findLength(A []int, B []int) int {
	n := len(A)
	m := len(B)
	max := 0
	var dp [1000][1000]int
	for i := 0; i < n; i++ {
		if A[i] == B[0] {
			dp[i][0] = 1
		}
	}
	for i := 0; i < m; i++ {
		if A[0] == B[i] {
			dp[0][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if A[i] == B[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}

	return max
}

func stoneGame(piles []int) bool {
	var dp [500][500]int
	n := len(piles)
	if n <= 2 {
		return true
	}
	for i := 0; i < n; i++ {
		dp[i][i] = piles[i]
	}
	for i := 1; i < n; i++ {
		dp[i-1][i] = max(piles[i-1], piles[i]) - min(piles[i-1], piles[i])
	}
	for m := 2; m < n; m++ {
		for i := 0; i < n-m; i++ {
			j := i + m
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	if dp[0][n-1] > 0 {
		return true
	} else {
		return false
	}
}

func stoneGameII(piles []int) int {
	n := len(piles)
	var dp [101][101]int
	sum := 0
	for i := n - 1; i >= 0; i++ {
		sum += piles[i]
		for j := 0; j <= n; j++ {
			if i+2*j >= n {
				dp[i][j] = sum
			} else {
				for x := 1; i+x < n && x <= 2*j; x++ {
					dp[i][j] = max(dp[i][j], sum-dp[i+x][max(j, x)])
				}
			}
		}
	}
	return dp[0][1]
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	var mapPrice [101][101]int
	for _, value := range flights {
		if value[0] == src {
			mapPrice[0][value[1]] = value[2]
		}
	}
	for k := 1; k <= K; k++ {
		for i := 0; i < n; i++ {
			mapPrice[k][i] = mapPrice[k-1][i]
		}
		for _, value := range flights {
			if mapPrice[k-1][value[0]] != 0 && value[1] != src {
				if mapPrice[k][value[1]] == 0 {
					mapPrice[k][value[1]] = mapPrice[k-1][value[0]] + value[2]
				} else {
					mapPrice[k][value[1]] = min(mapPrice[k-1][value[0]]+value[2], mapPrice[k][value[1]])
				}
			}
		}
	}
	if mapPrice[K][dst] != 0 {
		return mapPrice[K][dst]
	} else {
		return -1
	}
}

func rob1(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	pre, cur := 0, nums[0]
	for i := 1; i < n; i++ {
		temp := cur
		cur = max(cur, pre+nums[i])
		pre = temp
	}
	return cur
}

func maxTurbulenceSize(A []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	var trend bool
	cur := 1
	total := 1
	for ; cur < n; cur++ {
		if A[cur-1] < A[cur] {
			trend = true
			total += 1
			break
		} else if A[cur-1] > A[cur] {
			trend = false
			total += 1
			break
		}
	}
	maxNumber := total
	for i := cur + 1; i < n; i++ {
		if trend && A[i] < A[i-1] {
			total += 1
			trend = false
		} else if !trend && A[i] > A[i-1] {
			total += 1
			trend = true
		} else if A[i] == A[i-1] {
			maxNumber = max(maxNumber, total)
			total = 1
		} else {
			maxNumber = max(maxNumber, total)
			total = 2
		}
	}
	return max(maxNumber, total)
}

func numTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	wordMap := make(map[byte]int)
	length := 0
	maxLen := 0
	begin := 0
	for i := 0; i < n; i++ {
		if _, ok := wordMap[s[i]]; ok {
			if wordMap[s[i]] >= begin {
				maxLen = max(maxLen, length)
				length = i - wordMap[s[i]]
				begin = wordMap[s[i]]
			} else {
				length += 1
			}
			wordMap[s[i]] = i
		} else {
			wordMap[s[i]] = i
			length += 1
		}
	}
	maxLen = max(maxLen, length)
	return maxLen
}

func maxSubArray(nums []int) int {
	n := len(nums)
	cur := 0
	maxNumber := ^int(^uint(0) >> 1)
	for i := 0; i < n; i++ {
		cur = max(nums[i], nums[i]+cur)
		maxNumber = max(maxNumber, cur)
	}
	return maxNumber
}

func numSquares1(n int) int {
	dp := make([]int, n+1)
	var cur int
	var minNum int
	dp[0] = 0
	for i := 1; i <= n; i++ {
		cur = 1
		minNum = i + 1
		for cur*cur <= i {
			minNum = min(minNum, dp[i-cur*cur]+1)
			cur = cur + 1
		}
		dp[i] = minNum
	}
	return dp[n]
}

func integerBreak(n int) int {
	var result int
	return result
}

//func longestSubstring(s string, k int) int {
//
//}

func maxProduct(nums []int) int {
	maxNum, minNum, result := 1, 1, math.MinInt64
	for _, num := range nums {
		if num < 0 {
			maxNum, minNum = minNum, maxNum
		}
		maxNum, minNum = max(maxNum, maxNum*num), min(minNum, minNum*num)
		result = max(maxNum, result)
	}
	return result
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, 0)
	dp = append(dp, nums[0])
	for i := 1; i < n; i++ {
		left, right := 0, len(dp)
		for left < right {
			mid := left + (right-left)/2
			if dp[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[left] = nums[i]
		}
	}
	return len(dp)
}

func partMaxPathSum(root *TreeNode, hashMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if value, ok := hashMap[root]; ok {
		return value
	}
	cur := max(partMaxPathSum(root.Left, hashMap), partMaxPathSum(root.Right, hashMap)) + root.Val
	hashMap[root] = cur
	return cur
}

func maxPathSum(root *TreeNode) int {
	hashMap := make(map[*TreeNode]int)
	return partMaxPathSum(root, hashMap)
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

//func longestSubstring(s string, k int) int {
//
//}
