/*
* @Author: qiuyu
* @Date:   2019-04-26 03:24:05
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:02
*/
func isPalindrome(x int) bool {
    	if x < 0 {
		return false
	} else {
		temp := strconv.Itoa(x)
		p := []byte(temp)
		status := true
		for i := 0; i < len(p)/2; i++ {
			if p[i] != p[len(p)-i-1] {
				status = false
			}
		}
		return status
	}
}
//56 ms