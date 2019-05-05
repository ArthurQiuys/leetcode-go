/*
* @Author: qiuyu
* @Date:   2019-05-05 17:08:46
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:57:18
*/
package main

func main() {
	
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p|| root == q{
		return root
	}  
	left:= lowestCommonAncestor(root.Left,p,q)
	if left!=nil && left!= p && left!= q{
		return left
	}
	right := lowestCommonAncestor(root.Right,p,q)
	if left!=nil && right!=nil{
		return root
	}
	if left == nil{
		return right
	}
	return left

}