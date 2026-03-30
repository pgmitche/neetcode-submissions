/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    // guard
	if lists == nil || len(lists) == 0 {
		return nil
	}

	// iterate base case
	for len(lists) > 1 {
		tempMerged := []*ListNode{}

		// pair
		for i:=0; i< len(lists); i+=2 {
			var rhs *ListNode
			// guard rhs cursor out of bounds
			if i+1 < len(lists) {
				rhs = lists[i+1]
			}
			// append the merged list to the list
			tempMerged = append(tempMerged, mergeLists(lists[i], rhs))
		}

		// refine source list
		lists = tempMerged
	}

	return lists[0]
}

func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: how to sort linked list???
	// create dummy head
	dummy := &ListNode{}
	// mark tail as dummy
	tail := dummy

	// while there are values in both
	for list1 != nil && list2 != nil {
		// append the lowest value to the chain then
		// shrink the changed chain to advance it's next value
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		// advance the tail
		tail = tail.Next
	}

	// when there's a remaining node without a 
	// comparator
	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}

	// Skip the empty dummy!!
	return dummy.Next 
}