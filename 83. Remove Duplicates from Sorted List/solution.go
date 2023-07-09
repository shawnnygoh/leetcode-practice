/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    // Base case
    if head == nil || head.Next == nil {
        return head
    // Duplicate element
    } else if head.Val == head.Next.Val {
        head = deleteDuplicates(head.Next)
        return head
    } else {
        head.Next = deleteDuplicates(head.Next)
        return head
    }
}