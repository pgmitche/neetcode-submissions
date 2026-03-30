/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	// Perform a postordered dfs with height/balance evaluation
	var dfs func(*TreeNode) (bool, int)
	dfs = func(node *TreeNode) (bool, int) {
		// reach the base case, this is obviously balanced
		if node == nil {
			return true, 0
		}

		// traverse subtree
		lhsBalanced, lhsHeight := dfs(node.Left)
		rhsBalanced, rhsHeight := dfs(node.Right)

		// check if subtree is balanced
		balanced := lhsBalanced && rhsBalanced
		var heightDiff int
		if lhsHeight > rhsHeight {
			heightDiff = lhsHeight - rhsHeight
		} else {
			heightDiff = rhsHeight - lhsHeight
		}

		isBalanced := balanced && heightDiff <= 1

		// now track this nodes' height by incrementing from child's max
		thisHeight := 1 + max(lhsHeight, rhsHeight)
		return isBalanced, thisHeight
	}
	
	isBalanced, _ := dfs(root)
	return isBalanced
}
