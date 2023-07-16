/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Main Idea: Use 3 pointers `curr`, `prev` and `next` to keep track of nodes to 
// update reverse links.

// Iterative Solution
// Time Complexity: O(n), where n is the length of the linked list
// Space Complexity: O(1)

func reverseList(head *ListNode) *ListNode {
    // Create a variable `curr` and assign `head` to it.
    curr := head
    
    // If `curr` is empty, return it (which should be an empty list).
    if curr == nil {
        return curr
    
    // Else, declare a variable `prev` of type ListNode.
    } else {
        var prev *ListNode

        // `for` is Go's version of `while`.
        // While curr exists, 
        // 1. Set `next` to the next ListNode of curr
        // 2. Set the next ListNode of `curr` to `prev` (Arrow now points backwards)
        // 3. Set `prev` to `curr`
        // 4. Set `curr` to `next`
        for curr != nil {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }

        // Return `prev` which should contain the reversed linked list. 
        return prev
    }
}