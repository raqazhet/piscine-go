package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func hello(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = hello(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = hello(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
