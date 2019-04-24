/*
* @Author: qiuyu
* @Date:   2019-04-25 02:00:43
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-25 02:03:10
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