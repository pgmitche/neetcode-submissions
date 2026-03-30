/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    
	for list2 != nil && list1 != nil{
		if list1.Val < list2.Val {
			// add to dummy chain
			current.Next = list1
			// pop from front of list
			list1 = list1.Next
		} else {
			current.Next = list2
            list2 = list2.Next
		}
		// move tail pointer along
		current = current.Next 
	}

    // Append any remaining nodes from either list
	 if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

	// trim empty dummy value from front and return true
	return dummy.Next
}	
