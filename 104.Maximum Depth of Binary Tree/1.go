/*
* @Author: qiuyu
* @Date:   2019-04-29 23:15:18
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 23:21:23
*//**
 * /Definition for a binary tree node.
 */ 

package main
import (
	"math"
)
 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func main() {
	
}
func maxDepth(root *TreeNode) int {
    if root ==nil{
    	return 0
    }
    return 1 + math.MaxInt32(maxDepth(root.Left),maxDepth(root.Right))
}
func Max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}