/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    small := &ListNode{}
    smallHead := small // 记录头指针位置
    large := &ListNode{}
    largeHead := large
    // 将原链表分隔成两个链表
    for head != nil {
        if head.Val < x {
            small.Next = head
            small = small.Next
        } else {
            large.Next = head
            large = large.Next
        }
        // 移动原链表指针
        head = head.Next
    }
    // 连接两个链表
    large.Next = nil // 截断链表，否则会一直补0
    small.Next = largeHead.Next // 拼接large链表除头部空指针后的部分
    return smallHead.Next  // 返回链表除头部空指针的部分
}
