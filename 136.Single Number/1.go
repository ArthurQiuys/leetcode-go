/*
* @Author: qiuyu
* @Date:   2019-05-01 23:13:25
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-02 01:53:21
*/
package main

func main() {
	
}
func singleNumber(nums []int) int {
    res := 0
    for _,i:= range nums{
    	res ^= i
    }
    return res
}