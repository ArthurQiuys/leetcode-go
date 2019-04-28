/*
* @Author: qiuyu
* @Date:   2019-04-27 02:14:53
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:09
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
        return nil
    }
    n := len(lists)
    for n >1{
        k := (n+1)/2
        for i:= 0; i< n/2 ; i++{
            lists[i] = mergeTwoLists(lists[i],lists[i+k])
        }
        n = k;
    }
    return lists[0]
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode{
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