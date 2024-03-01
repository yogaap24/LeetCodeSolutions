/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    
    for current != nil && current.Next != nil {
        if current.Val != current.Next.Val {
            current = current.Next
        } else {
            current.Next = current.Next.Next
        }
    }
    
    return head
}