/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	inOrderLookup := make(map[int]int, 0)
	for i, v := range inorder {
		inOrderLookup[v] = i
	}

	cursor := 0
	var reconstruct func(int, int) *TreeNode
	reconstruct = func(lhs int, rhs int) *TreeNode {
		if lhs > rhs {
			return nil
		}

		// Current root value from preorder
        rootVal := preorder[cursor]
		// find preOrder[0] in inOrder, to split the tree
		// on the current node.
		midIndex := inOrderLookup[rootVal]
        
        // Create root node
        root := &TreeNode{Val: rootVal} // don't use cursor lookup after increment
        cursor++

		// lhs pointer for the inorder val:
		// build the left tree using the lhs val and the left
		// side of the middle index (this node)
		root.Left = reconstruct(lhs, midIndex-1)

		// rhs pointer for the inorder val:
		// build the right tree using the rhs val and the right
		// side of the middle index (this node)
		root.Right = reconstruct(midIndex+1, rhs)

		return root
	}

	// start from the whole range, the cursor finds
	// the root node with it's initial zero value lookup
	// in the preorder slice
	return reconstruct(0, len(inorder)-1)
}
