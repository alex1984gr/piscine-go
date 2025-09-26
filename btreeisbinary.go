package student

func BTreeIsBinary(root *TreeNode) bool {
	return helper(root, "", "")
}

func helper(node *TreeNode, minValue, maxValue string) bool {
	if node == nil {
		return true
	}
	if minValue != "" && node.Data <= minValue {
		return false
	}
	if maxValue != "" && node.Data >= maxValue {
		return false
	}
	return helper(node.Left, minValue, node.Data) &&
		helper(node.Right, node.Data, maxValue)
}
