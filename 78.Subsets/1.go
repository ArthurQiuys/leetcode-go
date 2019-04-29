/*
* @Author: qiuyu
* @Date:   2019-04-29 18:42:57
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 19:05:09
*/
package main
import (
	"sort"
)
func main() {
	nums := []int{1,2,3}
	nums1 :=subsets(nums)
	println(nums1)
}
func subsets(nums []int) [][]int {
    res:= make([][]int, 0)
    sort.Ints(nums)
    max := len(nums)
    println(len(nums))
    for i := 0; i < max; i++ {
    	println("nums is ",nums[i])
    	out := convertinttoset(nums,i)
    	res = append(res,out)
    }
    return res
}
func convertinttoset(nums [] int, t int) []int {
	sub := []int{}
	idx := 0
	for i := t; i > 0; i>>=1 {
		println(nums[i])
		if (i & 1) ==1{
			println("i 等于",i)
			sub =append(sub, nums[idx] )
			println("sub等于",nums[idx])
		}
		idx ++
	}
	return sub
}