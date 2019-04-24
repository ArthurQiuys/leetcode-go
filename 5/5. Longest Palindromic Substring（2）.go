/*
* @Author: qiuyu
* @Date:   2019-04-25 01:59:16
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-25 02:03:08
*/
func longestPalindrome(s string) string {
    var rest string
    dp := make([][] bool,len(s))
    for i := range dp{
        dp[i] = make([]bool, len(s))
        dp[i][i] = true
    }
  for j:=0; j<len(s); j++ {
    for i:=j; i>=0; i-- {
      dp[i][j] = (s[j] == s[i]) && (i>=j-1 || dp[i+1][j-1])
      if dp[i][j] && j-i+1 >= len(rest) {
        rest = s[i:j+1]
      }
    }
  }
    return rest
}
// 执行用时 : 76 ms, 在Longest Palindromic Substring的Go提交中击败了40.87% 的用户
// 内存消耗 : 6.6 MB, 在Longest Palindromic Substring的Go提交中击败了20.57% 的用户
