/*
* @Author: qiuyu
* @Date:   2019-04-23 02:41:43
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-23 02:53:41
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry_num := 0
    head := new(ListNode)
    cur := head
    for l1 != nil || l2 != nil || carry_num != 0{
        n1,n2 := 0, 0
        if l1 != nil{
            n1, l1 = l1.Val , l1.Next
        }
        if l2 != nil{
            n2, l2 = l2.Val , l2.Next
        }
        num := n1 + n2 + carry_num
        carry_num = num / 10
        cur.Next = &ListNode{Val:num%10, Next:nil}
        cur = cur.Next
    }
    return head.Next
}
//执行用时 : 40 ms, 在Add Two Numbers的Go提交中击败了25.00% 的用户
//内存消耗 : 5 MB, 在Add Two Numbers的Go提交中击败了82.44% 的用户