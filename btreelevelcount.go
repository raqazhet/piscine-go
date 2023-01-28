package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rDepth := BTreeLevelCount(root.Right)
	lDepth := BTreeLevelCount(root.Left)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func hi(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = hi(root.Left, data)
	} else {
		root.Right = hi(root.Right, data)
	}
	return root
}
