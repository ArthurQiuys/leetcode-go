/*
* @Author: qiuyu
* @Date:   2019-04-29 23:47:38
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-30 00:07:02
*/
package main
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func main() {
	
}

func maxPathSum(root *TreeNode) int {
    res := 0
    maxpat(root, &res)
    return res
}
func max(a int, b int) int {
	if a > b{
		return a
	}
	return b
}
func maxpat(node *TreeNode, res *int) int{
	if node == nil{
		return 0
	}
	left := max(maxpat(node.Left,res), 0)
	right := max(maxpat(node.Right,res), 0)
	res = max(res, left + right + node.Val)
	return max(left, right) + node.Val
}