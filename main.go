package palindrome_linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	// initialize slow and fast pointers
	slow := head
	fast := head

	// initialize stack
	stack := []int{}

	// as long as fast isn't nil, and fast.next isn't nil
	for fast != nil && fast.Next != nil {
		// append slow value to stack
		stack = append(stack, slow.Val)
		// advance both; remember they started at the same place, fast moves faster
		slow = slow.Next
		fast = fast.Next.Next
	}

	// as long as slow isn't nil
	for slow != nil {
		// is the slow value and the last element in the stack the same?
		if stack[len(stack)-1] == slow.Val {
			// pop the last node
			stack = stack[:len(stack)-1]
		}
		// move slow
		slow = slow.Next
	}

	// is it empty!?
	// if its empty then we have a palindrome
	return len(stack) == 0
}
