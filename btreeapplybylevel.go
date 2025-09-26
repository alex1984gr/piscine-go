package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	node := queue[0]
	queue = queue[1:]
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
	}
	if node.Left != nil {
		queue = append(queue, node.Left)
	}
	if node.Right != nil {
		queue = append(queue, node.Right)
	}
}
