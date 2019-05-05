/*
* @Author: qiuyu
* @Date:   2019-05-05 17:38:09
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:57:18
*/
package main

func main() {
	
}
func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    res [0] =1
    for i := 1; i < len(nums); i++ {
    	res[i] = res[i - 1] *nums[i-1]
    }
    right:=1
    for i := len(nums)-1; i >=0 ; i-- {
    	res[i] *= right
    	right *= nums[i]
    }
    return res
}
//超时
