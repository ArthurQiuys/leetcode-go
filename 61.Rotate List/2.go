/*
* @Author: qiuyu
* @Date:   2019-04-29 17:57:21
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 17:57:31
*/
package main

func main() {
	
}
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil|| k == 0{
        return head
    }
    n:=1
    cur := head
    for cur.Next != nil{
        n++
        cur = cur.Next
    } 
    cur.Next = head
    k %= n
    for i := 0; i < n-k; i++ {
    	cur = cur.Next
    }
    newhead := cur.Next
    cur.Next = nil
    return newhead
    
}