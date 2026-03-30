/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    // find the node
	if root == nil {
		return nil
	}

	// recursively call deleteNode to "search" for the val to delete
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		// does the node have 0, 1, or 2 children?
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left 
		} else {
			// if there are 2 children, then we can either choose
			// the min from the right, or the max from the left
			max := findMax(root.Left)
			root.Val = max.Val
			// now overwrite the subtree of the lhs max value
			// shifting the tree up 
			root.Left = deleteNode(root.Left, max.Val)
		}
	}

	return root
}

func findMax(root *TreeNode) *TreeNode {
	current := root
	for current != nil && current.Right != nil {
		current = current.Right
	}

	return current
}