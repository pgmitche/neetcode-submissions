/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		//maintain link
		nextTemp := current.Next

		// reverse
		current.Next = prev
		
		// shift
		prev = current
		current = nextTemp
	}

	return prev
}
