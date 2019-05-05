/*
* @Author: qiuyu
* @Date:   2019-05-05 18:35:16
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:39:35
*/
package main

func main() {
	
}
func reverseString(s []byte)  {
    left, right := 0, len(s)-1
    for left <right{
    	t:= s[left]
    	s[left] = s[right]
    	s[right] = t
    }
    return s
}