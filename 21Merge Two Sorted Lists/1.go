/*
* @Author: qiuyu
* @Date:   2019-04-27 02:13:51
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:10
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil{
        return l2
    }
    if l2 == nil{
        return l1
    }
    var head *ListNode
    if l1.Val <=l2.Val{
        l1.Next =mergeTwoLists(l1.Next, l2)
        return l1;
    }else{
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2;
    }
    return head
}