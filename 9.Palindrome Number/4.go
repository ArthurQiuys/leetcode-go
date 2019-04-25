/*
* @Author: qiuyu
* @Date:   2019-04-26 03:24:28
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-26 03:25:01
*/
func isPalindrome(x int) bool {
	var s int
	for X := x; X > 0; X /= 10 {
		y := X % 10
		s = s*10 + y
	}
	return (s == x)
}
//44ms
