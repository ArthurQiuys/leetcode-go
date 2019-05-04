/*
* @Author: qiuyu
* @Date:   2019-05-03 23:17:22
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-04 18:06:56
*/
package main

func main() {
	
}

func sortList(head *ListNode) *ListNode {
    if head ==nil || head.Next == nil{
    	return head
    }
    slow , fast:= head , head
    var pre *ListNode
    for fast != nil && fast.Next!= nil{
        pre = slow
    	slow = slow.Next
    	fast = fast.Next.Next
    }
    pre.Next = nil
    return merget(sortList(head),sortList(slow))
}

func merget(head *ListNode, slow *ListNode) *ListNode {
	if head == nil{
		return slow
	}
	if slow == nil{
		return head
	}
	if head.Val < slow.Val{
		head.Next = merget(head.Next,slow)
		return head
	}else{
		slow.Next = merget(head, slow.Next)
		return slow
	}
	return head

}