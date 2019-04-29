/*
* @Author: qiuyu
* @Date:   2019-04-28 21:29:41
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 17:57:34
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func main() {
	
}
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil|| k == 0{
        return head
    }
    n:=0
    cur := head
    for cur.Next != nil{
        n++
        cur = cur.Next
    } 
    cur.Next = head
    m:= n-k%n
    for i := 0; i < m; i++ {
    	cur = cur.Next
    }
    newhead := cur.Next
    cur.Next = nil
    return newhead
    
}