/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{}
    p := head
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            p.Next = list1
            p = list1
            list1 = list1.Next
        } else {
            p.Next = list2
            p = list2
            list2 = list2.Next
        }
    }
    if list1 == nil {
        p.Next = list2
    }
    if list2 == nil {
        p.Next = list1
    }
    return head.Next
}
