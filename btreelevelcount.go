package student

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
