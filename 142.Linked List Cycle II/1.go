/*
* @Author: qiuyu
* @Date:   2019-05-02 02:17:55
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-04 18:07:32
*/
package main

func main() {
	
}
func detectCycle(head *ListNode) *ListNode {
 	   fast , slow := head,head
 	   for fast != nil &&fast.Next !=nil{
 		slow = slow.Next
 		fast = fast.Next.Next
 		if fast == slow {
 			break
 		}
       }
 	if fast == nil || fast.Next ==nil{
 		return nil
 	}
    slow = head
 	for slow != fast{
 		slow = slow.Next
 		fast = fast.Next
 	}
    
 	return fast
}