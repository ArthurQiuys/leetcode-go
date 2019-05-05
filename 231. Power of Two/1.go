/*
* @Author: qiuyu
* @Date:   2019-05-05 16:50:41
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 16:54:06
*/
package main

func main() {
	
}
func isPowerOfTwo(n int) bool {
    if n > 0 && (n&(n-1)) == 0{
    	return true
    }
    return false
}