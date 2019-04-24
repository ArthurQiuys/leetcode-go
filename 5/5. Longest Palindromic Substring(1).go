/*
* @Author: qiuyu
* @Date:   2019-04-24 17:35:32
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-24 17:35:58
*/
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    max := string(s[0])
    n := len(s)
    for i := 0 ; i < n - 1; i++{
       max = searchPalindrome(s, i, i, max)
       max = searchPalindrome(s, i, i+1, max)
    }
    return max
}

func searchPalindrome(s string, i int, j int, max string) string {
    var sub string
    for i>=0 && j<len(s) && s[i] == s[j]{
        sub = s[i:j+1]
        i--;
        j++;
    }
    if len(max) < len(sub){
        max = sub
    }
    return max
}
/*执行用时 : 24 ms, 在Longest Palindromic Substring的Go提交中击败了55.15% 的用户
内存消耗 : 2.3 MB, 在Longest Palindromic Substring的Go提交中击败了69.86% 的用户*/
