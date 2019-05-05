/*
* @Author: qiuyu
* @Date:   2019-05-05 14:57:15
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 16:48:57
*/
package main

func main() {
	
} 
type TreeNode struct { 
     Val int
     Left *TreeNode
     Right *TreeNode
  }

func kthSmallest(root *TreeNode, k int) int {
	if root ==nil{
		return 0
	}
	cut := cont(root.Left)
	if k <= cut{
		return kthSmallest(root.Left,k)
	}else if k > cut + 1{
		return kthSmallest(root.Right, k-cut -1)
	}
	return root.Val
}
func cont(root *TreeNode) int{
	if root == nil{
		return 0
	}
	return 1 + cont(root.Right) +cont (root.Left)
}