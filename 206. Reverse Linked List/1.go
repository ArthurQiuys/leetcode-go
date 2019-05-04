/*
* @Author: qiuyu
* @Date:   2019-05-04 21:10:32
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-04 21:18:35
*/
package main
type ListNode struct {
     Val int
    Next *ListNode
  }
func main() {
	
}
func reverseList(head *ListNode) *ListNode {
    var newHead *ListNode 
    for head != nil{
    	t := head.Next
    	head.Next = newHead
    	newHead = head
    	head = t
    }
    return newHead
}
