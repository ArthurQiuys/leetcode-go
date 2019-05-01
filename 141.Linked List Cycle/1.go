/*
* @Author: qiuyu
* @Date:   2019-05-02 01:54:01
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-02 02:04:21
*/
package main

func main() {
	
}
func hasCycle(head *ListNode) bool {
 	fast , slow := head , head
 	for fast != nil &&fast.Next !=nil{
 		slow = slow.Next
 		fast = fast.Next.Next
 		if fast == slow {
 			return true
 		}

 	}   
 	return false
}