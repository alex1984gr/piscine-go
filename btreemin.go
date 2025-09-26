package student

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	minNode := root

	leftMin := BTreeMin(root.Left)
	if leftMin != nil && leftMin.Data < minNode.Data {
		minNode = leftMin
	}
	rightMin := BTreeMin(root.right)
	if rightMin != nil && rightMin.Data < minNode.Data {
		minNode = rightMin
	}
	return minNode
}
