package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	trees := make([]*TreeNode, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		result = append(result, cur.Val)
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			result = append(result, cur.Val)
		}
		if cur.Right != nil {
			cur = cur.Right
			trees = append(trees, cur)
		}

	}
	return result
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	trees := make([]*TreeNode, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		if cur.Right != nil {
			trees = append(trees, cur.Right)
		}
		if cur.Left != nil {
			trees = append(trees, cur.Left)
		}
		result = append(result, cur.Val)
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	treeFirst := make([]*TreeNode, 0)
	treeSecond := make([]*TreeNode, 0)
	treeFirst = append(treeFirst, cur)
	for len(treeFirst) != 0 {
		cur = treeFirst[len(treeFirst)-1]
		treeFirst = treeFirst[:len(treeFirst)-1]
		treeSecond = append(treeSecond, cur)
		if cur.Left != nil {
			treeFirst = append(treeFirst, cur.Left)
		}
		if cur.Right != nil {
			treeFirst = append(treeFirst, cur.Right)
		}
	}
	for len(treeSecond) != 0 {
		cur = treeSecond[len(treeSecond)-1]
		treeSecond = treeSecond[:len(treeSecond)-1]
		result = append(result, cur.Val)
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	cur := root
	end := root
	nextEnd := root
	trees := make([]*TreeNode, 0)
	partResult := make([]int, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append([]*TreeNode{cur.Left}, trees...)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append([]*TreeNode{cur.Right}, trees...)
		}
		if cur == end {
			end = nextEnd
			result = append(result, partResult)
			partResult = partResult[0:0]
		}
	}
	return result
}

func partBalance(root *TreeNode, height int) (int, bool) {
	if root == nil {
		return height, true
	}
	leftHeight, leftBalance := partBalance(root.Left, height+1)
	rightHeight, rightBalance := partBalance(root.Right, height+1)
	ifBalance := leftBalance && rightBalance
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		ifBalance = false
	}
	return max(rightHeight, leftHeight), ifBalance
}

func isBalanced1(root *TreeNode) bool {
	_, ifBalance := partBalance(root, 0)
	return ifBalance
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	cur := root
	end := root
	nextEnd := root
	trees := make([]*TreeNode, 0)
	partResult := make([]int, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append([]*TreeNode{cur.Left}, trees...)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append([]*TreeNode{cur.Right}, trees...)
		}
		if cur == end {
			end = nextEnd
			result = append([][]int{partResult}, result...)
			partResult = make([]int, 0)
		}
	}
	return result
}

func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	if root == nil {
		return true
	}
	trees := make([]*TreeNode, 0)
	cur := root
	trees = append(trees, cur)
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		result = append(result, cur.Val)
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			result = append(result, cur.Val)
		}
		if cur.Right != nil {
			trees = append(trees, cur.Right)
			cur = cur.Right
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func kthSmallest2(root *TreeNode, k int) int {
	trees := make([]*TreeNode, 0)
	trees = append(trees, root)
	cur := root
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		if k > 1 {
			k -= 1
		} else {
			return cur.Val
		}
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			if k > 1 {
				k -= 1
			} else {
				return cur.Val
			}
		}
		if cur.Right != nil {
			cur = cur.Right
			trees = append(trees, cur)
		}
	}
	return -1
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	trees := make([]*TreeNode, 0)
	trees = append(trees, root)
	end, nextEnd := root, root
	partResult := make([]int, 0)
	earse := false
	for len(trees) != 0 {
		cur := trees[0]
		trees = trees[1:]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append(trees, cur.Left)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append(trees, cur.Right)
		}
		if cur == end {
			end = nextEnd
			if earse {
				partResult = reverse(partResult)
			}
			result = append(result, partResult)
			partResult = make([]int, 0)
			earse = !earse
		}
	}
	return result
}

func partBuildTree1(preorder []int, inMap map[int]int, left int, right int, index *int) *TreeNode {
	if left > right {
		return nil
	}
	newNode := new(TreeNode)
	number := preorder[*index]
	*index += 1
	cur := inMap[number]
	leftNode := partBuildTree1(preorder, inMap, left, cur-1, index)
	rightNode := partBuildTree1(preorder, inMap, cur+1, right, index)
	newNode.Val, newNode.Left, newNode.Right = number, leftNode, rightNode
	return newNode
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	inMap := make(map[int]int)
	for index, value := range inorder {
		inMap[value] = index
	}
	a := 0
	root := partBuildTree1(preorder, inMap, 0, n-1, &a)
	return root
}

func partBuildTree(postorder []int, inMap map[int]int, left int, right int, index *int) *TreeNode {
	if left > right {
		return nil
	}
	newNode := new(TreeNode)
	number := postorder[*index]
	*index -= 1
	cur := inMap[number]
	rightNode := partBuildTree(postorder, inMap, cur+1, right, index)
	leftNode := partBuildTree(postorder, inMap, left, cur-1, index)
	newNode.Val, newNode.Left, newNode.Right = number, leftNode, rightNode
	return newNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	inMap := make(map[int]int)
	for index, value := range inorder {
		inMap[value] = index
	}
	a := n - 1
	root := partBuildTree(postorder, inMap, 0, n-1, &a)
	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	cur := root
	temp := cur
	for {
		if cur == nil {
			break
		}
		if cur.Val < val {
			temp = cur
			cur = cur.Right
		} else {
			temp = cur
			cur = cur.Left
		}
	}
	newNode := new(TreeNode)
	newNode.Val, newNode.Left, newNode.Right = val, nil, nil
	if temp == nil {
		return newNode
	}
	if temp.Val > val {
		temp.Left = newNode
	} else {
		temp.Right = newNode
	}
	return root
}

func partRob(root *TreeNode, robMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _, ok := robMap[root]; ok {
		return robMap[root]
	}
	var robLeft, robRight int
	if root.Left == nil {
		robLeft = 0
	} else {
		robLeft = partRob(root.Left.Left, robMap) + partRob(root.Left.Right, robMap)
	}
	if root.Right == nil {
		robRight = 0
	} else {
		robRight = partRob(root.Right.Left, robMap) + partRob(root.Right.Right, robMap)
	}
	robValue := max(robRight+robLeft+root.Val, partRob(root.Left, robMap)+partRob(root.Right, robMap))
	robMap[root] = robValue
	return robValue
}

func rob(root *TreeNode) int {
	robMap := make(map[*TreeNode]int)
	return partRob(root, robMap)
}

func partGenerateTrees(begin int, end int) []*TreeNode {
	if begin > end {
		return []*TreeNode{nil}
	}

	result := make([]*TreeNode, 0)
	for i := begin; i <= end; i++ {
		leftNodes := partGenerateTrees(begin, i-1)
		rightNodes := partGenerateTrees(i+1, end)
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				newNode := new(TreeNode)
				newNode.Left, newNode.Right, newNode.Val = leftNode, rightNode, i
				result = append(result, newNode)
			}
		}
	}
	return result
}
func generateTrees(n int) []*TreeNode {
	return partGenerateTrees(1, n)
}

type Trie struct {
	children [26]*Trie
	end      bool
}

/** Initialize your data structure here. */
func Constructor2() Trie {
	var root Trie
	var children [26]*Trie
	root.children = children
	return root
}

func (this *Trie) partInsert(word string, index int, n int) {
	if index == n {
		return
	}

	if this.children[word[index]-'a'] == nil {
		newNode := new(Trie)
		var children [26]*Trie
		newNode.children = children
		this.children[word[index]-'a'] = newNode
	}

	if index == n-1 {
		this.children[word[index]-'a'].end = true
	}

	this.children[word[index]-'a'].partInsert(word, index+1, n)
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.partInsert(word, 0, len(word))
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		cur = cur.children[word[i]-'a']
		if cur == nil {
			return false
		}
	}
	if cur.end == true {
		return true
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		cur = cur.children[prefix[i]-'a']
		if cur == nil {
			return false
		}
	}
	return false
}

func rightSideView(root *TreeNode) []int {
	var last, nextLast, cur *TreeNode = root, nil, nil
	queue := make([]*TreeNode, 0)
	queue = append(queue, last)
	result := make([]int, 0)
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			nextLast = cur.Left
			queue = append(queue, nextLast)
		}
		if cur.Right != nil {
			nextLast = cur.Right
			queue = append(queue, nextLast)
		}
		if cur == last {
			last = nextLast
			result = append(result, cur.Val)
		}
	}
	return result
}

func pathTo1(root, p, q *TreeNode, parents **TreeNode) bool {
	if root == nil {
		return false
	}
	successL := pathTo1(root.Left, p, q, parents)
	successR := pathTo1(root.Right, p, q, parents)
	var part int = 0
	if successL {
		part += 1
	}
	if successR {
		part += 1
	}
	if root == q || root == p {
		part += 1
	}
	if part >= 2 {
		*parents = root
	}
	return part > 0
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode = nil
	pathTo1(root, p, q, &result)
	return result
}

func partInvertTree(root *TreeNode) {

}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	trees := make([]*TreeNode, 0)
	trees = append(trees, root)
	var cur *TreeNode
	for len(trees) != 0 {
		cur = trees[0]
		trees = trees[1:]
		cur.Left, cur.Right = cur.Right, cur.Left
		if cur.Left != nil {
			trees = append(trees, cur.Left)
		}
		if cur.Right != nil {
			trees = append(trees, cur.Right)
		}
	}
	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		if p == nil && q == nil {
			return true
		} else {
			return false
		}
	}
}

func partSymmetric(left *TreeNode, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		return partSymmetric(left.Left, right.Right) && partSymmetric(left.Right, right.Left)
	} else {
		if left == nil && right == nil {
			return true
		} else {
			return false
		}
	}

}

func isSymmetric(root *TreeNode) bool {
	return partSymmetric(root.Left, root.Right)
}

func partCommonAncestor(root, p, q *TreeNode, part **TreeNode) bool {
	if root == nil {
		return false
	}
	result := 0
	if partCommonAncestor(root.Left, p, q, part) {
		result += 1
	}
	if partCommonAncestor(root.Right, p, q, part) {
		result += 1
	}
	if root == p || root == q {
		result += 1
	}
	if result >= 2 {
		*part = root
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode
	partCommonAncestor(root, p, q, &result)
	return result
}

func partPathSum1(root *TreeNode, sum int, cur int) bool {
	if root == nil {
		return false
	}
	cur = cur + root.Val
	if cur == sum {
		return true
	} else if cur > sum {
		return false
	}
	return partPathSum1(root.Left, sum, cur) || partPathSum1(root.Left, sum, cur)
}

func hasPathSum1(root *TreeNode, sum int) bool {
	return partPathSum1(root, sum, 0)
}

func partPathSum(root *TreeNode, sum int, curList []int, cur int, result *[][]int) {
	if root == nil {
		return
	}
	cur = cur + root.Val
	if cur == sum && root.Left == nil && root.Right == nil {
		*result = append(*result, curList)
	}
	nowCurList := make([]int, len(curList))
	copy(nowCurList, curList)
	nowCurList = append(nowCurList, root.Val)
	partPathSum(root.Left, sum, curList, cur, result)
	partPathSum(root.Right, sum, curList, cur, result)
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	curList := make([]int, 0)
	partPathSum(root, sum, curList, 0, &result)
	return result
}

func partDiameterOfBinaryTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftHeight, leftMax := partDiameterOfBinaryTree(root.Left)
	rightHeight, rightMax := partDiameterOfBinaryTree(root.Right)
	return max(leftHeight, rightHeight) + 1, max(leftHeight+rightHeight+1, max(leftMax, rightMax))
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, result := partDiameterOfBinaryTree(root)
	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func partBalanceS(root *TreeNode, result *bool) int {
	if !*result {
		return 0
	}
	if root == nil {
		return 0
	}
	leftHeight := partBalanceS(root.Left, result)
	rightHeight := partBalanceS(root.Right, result)
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		*result = false
	}
	return max(leftHeight, rightHeight) + 1
}

func isBalanced(root *TreeNode) bool {
	result := true
	partBalanceS(root, &result)
	return result
}

func invertTree(root *TreeNode) *TreeNode {
	trees := make([]*TreeNode, 0)
	trees = append(trees, root)
	var cur *TreeNode
	for len(trees) != 0 {
		cur = trees[0]
		trees = trees[1:]
		cur.Right, cur.Left = cur.Left, cur.Right
		if cur.Left != nil {
			trees = append(trees, cur.Left)
		}
		if cur.Right != nil {
			trees = append(trees, cur.Right)
		}
	}
	return root
}

func partMergeTrees(t1 *TreeNode, t2 *TreeNode) {
	if t2 == nil || t1 == nil {
		return
	}
	if t1.Left == nil && t2.Left != nil {
		t1.Left = t2.Left
		t2.Left = nil
	}
	if t1.Right == nil && t2.Right != nil {
		t1.Right = t2.Right
		t2.Right = nil
	}
	t1.Val = t1.Val + t2.Val
	partMergeTrees(t1.Left, t2.Left)
	partMergeTrees(t1.Right, t2.Right)
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	partMergeTrees(t1, t2)
	return t1
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	trees, treeNum := make([]*TreeNode, 0), make([]int, 0)
	var maxNumber, curNum = 1, 0
	var end, nextEnd, cur *TreeNode
	trees, treeNum, end = append(trees, root), append(treeNum, 0), root
	for len(trees) != 0 {
		cur, curNum, trees, treeNum = trees[0], treeNum[0], trees[1:], treeNum[1:]
		if cur.Left != nil {
			trees, nextEnd = append(trees, cur.Left), cur.Left
			treeNum = append(treeNum, curNum*2)
		}
		if cur.Right != nil {
			trees, nextEnd = append(trees, cur.Right), cur.Right
			treeNum = append(treeNum, curNum*2+1)
		}
		if cur == end && len(trees) != 0 {
			end, maxNumber = nextEnd, max(maxNumber, treeNum[len(treeNum)-1]-treeNum[0]+1)
		}
	}
	return maxNumber
}
