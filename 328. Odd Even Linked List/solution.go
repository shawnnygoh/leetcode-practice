/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Main Idea: Split the linked list into two lists which contain the odd and even nodes respectively using pointers. Then, append the linked list containing even nodes at the end of the linked list containing odd nodes.
// Time Complexity: O(n), where n is the length of the linked list since the linked list is traversed only once.
// Space Complexity: O(1)

func oddEvenList(head *ListNode) *ListNode {
    if head != nil {
        odd := head
        even := head.Next
        evenStart := even

        // While not at the end of the list,
        for odd.Next != nil && even.Next != nil {

            // Linking all odd nodes
            odd.Next = even.Next

            // Assigning the next odd node to `odd`
            odd = odd.Next

            // Linking all even nodes
            even.Next = odd.Next

            // Assigning the next even node to `even`
            even = even.Next
        }
        // Append the even node linked list to the odd node linked list by placing the start of the even node 
        // linked list (the first even node) after the end of the odd node linked list (the last odd node).
        odd.Next = evenStart
    }

    // Return odd even linked list.
    return head
}