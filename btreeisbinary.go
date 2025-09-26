package student

func BTreeIsBinary(root *TreeNode) bool {
	return helper(root, "", "")
}
func helper(node *TreeNode, minValeu, maxValeu string) bool {
	if node == nil {
		return true
	}
	if minValeu != "" && node.Data <= minValeu {
		return false
	}
	if maxValeu != "" && node.Data >= maxValeu {
		return false
	}
	return helper(node.Left, minValeu, node.Data) &&
		helper(node.Right, node.Data, maxValeu)
}
