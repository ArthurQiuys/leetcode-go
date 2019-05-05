/*
* @Author: qiuyu
* @Date:   2019-05-04 21:04:27
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 14:13:50
*/
package main

func main() {
	
}
func majorityElement(nums []int) int {
    res ,cut := 0 ,0
    for _ ,num := range nums{
    	if cut == 0{
    		res = num 
    		cut ++
    	}else if num ==res{
    		cut ++
    	}else if num !=res{
    		cut --
    	}
    }
    return res
}