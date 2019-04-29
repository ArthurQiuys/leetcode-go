/*
* @Author: qiuyu
* @Date:   2019-04-29 18:38:15
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 18:40:41
*/
package main

func main() {
	
}
func climbStairs(n int) int {
    if n == 1{
    	return n
    }
    nums := make([]int, n)
    nums[0] = 1
    nums[1] = 2
    for i := 2; i < n; i++ {
    	nums[i] = nums[i-1] +nums[i-2]
    }
    return nums[n-1]
}