/*
* @Author: qiuyu
* @Date:   2019-04-26 03:22:22
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:03
*/
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    sign := 1
    for ;x/sign >= 10;{
        sign = sign *10
    }
    var left,right int
    for ;x > 0;{
        left = x / sign
        right = x % 10
        if left != right{
            return false
        }
        x = (x % sign) /10
        sign = sign / 100
    }
    return true
}
//72 ms