/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// if the root is now empty, return the new node!
    if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val > root.Val {
		// if the value is greater, traverse right	
		root.Right = insertIntoBST(root.Right, val)
	} else if val < root.Val {
		// if the value is lower, traverse left	
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}
