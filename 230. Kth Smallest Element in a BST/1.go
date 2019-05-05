/*
* @Author: qiuyu
* @Date:   2019-05-05 14:44:50
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 14:57:08
*/
package main

func main() {
	
}
func kthSmallest(root *TreeNode, k int) int {
    return kthSmallestDFS(root, &k)
}
func kthSmallestDFS(root *TreeNode, k *int) int{
	if root == nil{
		return -1
	}
	val := kthSmallestDFS(root.Right,k)
	if *k == 0{
		return val
	}
	if *k-1 == 0{
		*k-=1
		return root.Val
	}
	return kthSmallestDFS(root.Left, k)
}