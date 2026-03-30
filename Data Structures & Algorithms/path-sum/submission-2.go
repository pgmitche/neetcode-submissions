/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var tracker = &Tracker{stack: make([]int, 0)}
type Tracker struct {
	sum int
	stack []int
}

func (t *Tracker) backtrack () {
	// pop
	var val int
	val, t.stack = t.stack[len(t.stack)-1], t.stack[:len(t.stack)-1]
	// recalc
	t.sum -= val
}

func (t *Tracker) track (val int) {
	// push
	t.stack = append(t.stack, val)
	// calc
	t.sum += val
}

func hasPathSum(root *TreeNode, targetSum int) bool {
 	// Create tracker to handle sum and stack
    tracker := &Tracker{
        sum: 0,
        stack:      []int{},
    }

	var search func( *TreeNode) bool
	search = func(node *TreeNode) bool {
		// guard for empty
		if node == nil {
			return false
		}	

		tracker.track(node.Val)

		// if leaf node
		if node.Left == nil && nil == node.Right {
			// if we've reached the sum at a leaf, success
			if tracker.sum == targetSum {
				return true
			}
			// otherwise backtrack and fail
			tracker.backtrack()
			return false
		} 

		// search subtrees
		if search(node.Left) {
			return true
		}

		if search(node.Right) {
			return true
		}

		// backtrack the subtree
		tracker.backtrack()
		return false
	}

	return search(root)
}
