/*
* @Author: qiuyu
* @Date:   2019-05-04 21:18:58
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-04 21:23:16
*/
package main
  type ListNode struct {
      Val int
      Next *ListNode
  }
func main() {
	

}

func reverseList(head *ListNode) *ListNode {
    if head ==nil || head.Next == nil{
    	return head
    }
    var newHead *ListNode
    newHead = reverseList(head)
    head.Next.Next = head
    head.Next = nil
    return  newHead
}