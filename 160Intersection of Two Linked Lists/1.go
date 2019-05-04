/*
* @Author: qiuyu
* @Date:   2019-05-04 18:33:26
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-04 21:02:48
*/
package main

func main() {
	
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headB == nil || headA == nil{
    	return nil
    }
    lenA ,lenB := getLen(headA), getLen(headB)
    if lenA < lenB{
    	for i := 0; i < lenB - lenA; i++ {
    		headB = headB.Next
    	}
    }else {
    	for i := 0; i < lenA - lenB; i++ {
    		headA = headA.Next
    	}
    }
    for headA!=nil && headB!=nil && headA != headB{
    	headA = headA.Next
    	headB = headB.Next
    }
    return headA
}
func getLen(head *ListNode) int{
	cut := 0
	for head !=nil {
		cut ++
		head = head.Next
	}
	return cut
}