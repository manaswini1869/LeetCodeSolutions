/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Initialize dummy node
    dummy := &ListNode{0, nil}
    dummy.Next = head

    // Initialize pointers for two-pass algorithm
    left := dummy
    right := head

    // Move right pointer to the nth node from the start
    for i := 0; i < n; i++ {
        right = right.Next
    }

    // Move both pointers until right reaches the end
    for right != nil {
        left = left.Next
        right = right.Next
    }

    // Remove the nth node from the end
    if left.Next != nil {
        left.Next = left.Next.Next
    }

    // Return the head of the modified list
    return dummy.Next
}