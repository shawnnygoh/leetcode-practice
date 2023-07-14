/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Utilizes Floyd's Cycle Finding Algorithm / Hare-Tortoise Algorithm 
// Main Idea: Move two pointers at different speeds, if they meet at some point in
// time then there exists a loop in the linked list. 

// Time Complexity: O(n), where n is the length of the linked list
// Space Complexity: O(1)

func hasCycle(head *ListNode) bool {
    // If head is empty or has no Next value then there cannot be a cycle.
    if head == nil || head.Next == nil {
        return false
    // Else, declare and assign variables `slow` as head and `fast` as the Next of
    // head.
    } else {
        slow := head
        fast := head.Next

        // `for` loop ensures that the fast pointer can move forward by one or two
        // nodes without encountering a nil reference, and that the fast pointer has
        // not met to the slow pointer (indicating the absence of a cycle). 
        for fast.Next != nil && fast.Next.Next != nil && fast != slow {
            slow = slow.Next
            fast = fast.Next.Next
        }

        // If slow is equal to fast, the fast pointer has caught up to the slow 
        // pointer and there is a loop. Else, the fast pointer has reached the end
        // of the linked list without meeting the slow pointer and hence there is no 
        // loop. 
        return slow == fast 
    }
}