/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	// in order to avoid passing a pointer to the slice
	// use an anonymous inline func to do the recursion
	var inOrder func(node *TreeNode)
    inOrder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inOrder(node.Left)
        result = append(result, node.Val)
        inOrder(node.Right)
    }

	inOrder(root)
	return result
}
