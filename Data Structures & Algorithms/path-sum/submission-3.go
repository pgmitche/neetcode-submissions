/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Reduce the value search as you traverse, less intuitive than my initial
// tracker stack based approach but easy to reason about as you think
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    // Check if we're at a leaf node with matching sum
    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }
    
    // Reduce target by current node value and check subtrees
    remaining := targetSum - root.Val
    return hasPathSum(root.Left, remaining) || hasPathSum(root.Right, remaining)
}
