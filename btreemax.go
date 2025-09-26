package student

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	maxNode := root

	leftMax := BTreeMax(root.Left)
	if leftMax != nil && leftMax.Data > maxNode.Data {
		maxNode = leftMax
	}
	rightMax := BTreeMax(root.Right)
	if rightMax != nil && rightMax.Data > maxNode.Data {
		maxNode = rightMax
	}
	return maxNode
}
