package student

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	if current.Right != nil {
		current = current.Right
	}
	return current
}
