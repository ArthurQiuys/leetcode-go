/*
* @Author: qiuyu
* @Date:   2019-04-29 17:58:26
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 18:16:14
*/
package main

func main() {
	n:=3
	m :=2
	t:=uniquePaths(n, m)
	println(t)
}
func uniquePaths(m int, n int) int {

    dp := make([]int,n)
    for i := 0; i < n; i++ {
    	dp[i] = 1
    }
    for i := 1; i < m; i++ {
    	for j := 1; j < n; j++ {
    		dp[j] += dp[j-1]
    	}
    }
    return dp[n-1]

}