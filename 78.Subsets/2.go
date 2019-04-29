/*
* @Author: qiuyu
* @Date:   2019-04-29 19:21:24
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-30 00:14:55
*/
package main
import (
	"sort"
)
func main() {
	
}
func subsets(nums []int) [][]int {
    res := [][]int{[]int{}}
    sort.Ints(nums)
    for _ , i := range nums {
    	 lens := len(res)
    	// var res []int
    	for j := 0; j < lens; j++ {
    		temp := make([]int, len(res[j]))
    		copy(temp, res[j])
    		temp = append(temp, i)
    		res = append(res,temp)
    		// res = append(res,nums[i] )
    	}
    	// res = append(res,res1 )
    }
    return res
}