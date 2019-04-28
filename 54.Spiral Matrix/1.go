/*
* @Author: qiuyu
* @Date:   2019-04-28 19:52:38
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-28 21:23:06
*/
package main
import (
	"fmt"
)
func main() {
	matrix :=[][]int{}
	t:= [] int{0}
	t =spiralOrder(matrix)
	fmt.Println(t)

}
func spiralOrder(matrix [][]int) []int {
    res := []int{0}
    /* if matrix == nil||len(matrix) == 0{
    	 return res
     }*/
    top, left := 0,0
    bottom , right := len(matrix)-1, len(matrix[0])-1
    for top <= bottom && left<= right{
    	for i := left; i < right; i++ {
    		res = append(res,matrix[top][i] )
    	}
    	if top+1 > bottom{
    		top++
    		break
    	}
    	for i := top; i < bottom; i++ {
    		res = append(res, matrix[i][right] )
    	}
    	if right -1 <left{
    		right -- 
    		break
    	}
    	for i := right; i >= left; i-- {
    		res = append(res ,matrix[bottom][i] )
    	}
    	if (bottom -1 < top){
    		bottom -- 
    		break
    	}
    	for i := bottom; i >= top; i-- {
    		res = append(res ,matrix[i][left] )
    		
    	}
    	if left +1 > right{
    		left++
    		break
    	}
    }
    return res
}