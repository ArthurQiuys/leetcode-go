/*
* @Author: qiuyu
* @Date:   2019-05-05 18:54:54
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:56:50
*/
package main

func main() {
	
}
func reverseWords(s string) string {
	res := []byte(s)
	var start, end int
	length := len(s)
	for i := 0; i < length; i++ {
		if i == length-1 || res[i+1] == ' ' {
			end = i
			for start < end {
				res[start], res[end] = res[end], res[start]
				start++
				end--
			}
			i++
			start = i + 1
		}
	}
	return string(res)
}