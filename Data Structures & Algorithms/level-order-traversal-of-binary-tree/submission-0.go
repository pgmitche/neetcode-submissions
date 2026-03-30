/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    // create a queue
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	// create the response
	levels := make([][]int, 0)
	for len(queue) > 0 {
		rowSize := len(queue)
		row := make([]int, 0, rowSize)

		// Process all nodes at current level
        for i := 0; i < rowSize; i++ {
			// pop, add node value to level
			var node *TreeNode
			node, queue = queue[0], queue[1:]
			row = append(row, node.Val)
			
			// Add children for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// then append the built row
		levels = append(levels, row)
	}

	return levels
}
