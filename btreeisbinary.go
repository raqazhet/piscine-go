package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}

	return true
}

func h(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = h(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = h(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
