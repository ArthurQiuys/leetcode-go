/*
* @Author: qiuyu
* @Date:   2019-05-05 14:42:42
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 14:43:33
*/
package main

func main() {
	
}
func containsDuplicate(nums []int) bool {
    m := make(map[int]bool ,len(nums))
    for _,n:=range nums{
        if m[n] == true{
        	return true
        }else{
        	m[n] = true
        }
    }
    return false
}