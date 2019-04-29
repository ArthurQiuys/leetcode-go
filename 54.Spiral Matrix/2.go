/*
* @Author: qiuyu
* @Date:   2019-04-28 19:55:23
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 19:22:03
*/
package main

func main() {
	
}
func spiralOrder(matrix [][]int) []int {
	if matrix ==nil||len(matrix)==0{
		return nil
	}
	top , left:= 0,0
	bottom ,right := len(matrix)-1, len(matrix[0])-1
	res := make([]int, 0)
	for top<= bottom && left <=right{
		for i := left; i <= right; i++ {
			res = append(res,matrix[top][i] )
		}
		top ++
		for i := top; i <= bottom; i++ {
			res = append(res,matrix[i][right] )
		}
		right--
	
		if top <= bottom{
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i] )
			}
			bottom --
		}
		if left <= right{
			for i := bottom; i >= top; i-- {
				res = append(res,matrix[i][left] )
			}
		}
			left++
	}
	return res
}