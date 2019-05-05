/*
* @Author: qiuyu
* @Date:   2019-05-05 17:18:48
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 17:20:48
*/
package main

func main() {
	
}
func deleteNode(node *ListNode) {
    node.Val=node.Next.Val
    node.Next = node.Next.Next
}
