/*
* @Author: qiuyu
* @Date:   2019-04-26 03:23:54
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:03
*/
func isPalindrome(x int) bool {
    if x<0 ||( x% 10 == 0 && x!=0){
        return false
    }
    endnum := 0
    for x > endnum{
        endnum = endnum * 10 + x % 10
        x /= 10
    }
    if x == endnum || x == endnum / 10{
        return true
    }
    return false
}
//68 ms