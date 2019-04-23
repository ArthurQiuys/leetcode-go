/*
* @Author: qiuyu
* @Date:   2019-04-23 02:54:16
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-23 03:07:15
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if (l1 == nil){
        return l2;
    }
    if (l2 == nil){
        return l1;
    }
    value := l1.Val + l2.Val
    if (value >=10){
        value = value -10
        if(l1.Next == nil && l2.Next == nil){
            next := &ListNode{1, nil}
            l3:= &ListNode{Val: value, Next: next}
            return l3
        }
        if(l1.Next != nil && l2.Next != nil){
            l1.Next.Val = l1.Next.Val + 1
            next := addTwoNumbers(l1.Next, l2.Next)
            l3:= &ListNode{Val: value, Next: next}
            return l3
        }
        if (l1.Next == nil){
            l1.Next = &ListNode{1, nil}
        }
        if (l2.Next == nil){
            l2.Next = &ListNode{1, nil}
        }
        
    }
    next := addTwoNumbers(l1.Next, l2.Next)
    l3 := &ListNode{Val : value, Next: next}
    return l3
}
//执行用时 : 24 ms, 在Add Two Numbers的Go提交中击败了91.71% 的用户
//内存消耗 : 5.1 MB, 在Add Two Numbers的Go提交中击败了46.69% 的用户