/*
* @Author: qiuyu
* @Date:   2019-05-05 16:57:32
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 17:07:28
*/
package main

func main() {
	
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val > max(p.Val ,q.Val){
		return lowestCommonAncestor(root.Left,p,q)
	}else if root.Val < min(p.Val, q.Val){
		return lowestCommonAncestor(root.Right,p,q)
	}
	return root
}
func max(a int ,b int)int{
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int{
	if a> b{
		return b
	}
	return a
}