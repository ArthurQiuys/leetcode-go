/*
* @Author: qiuyu
* @Date:   2019-04-29 23:43:48
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 23:45:50
*/
package main

func main() {
	
}
func maxProfit(prices []int) int {
    res , n := 0 ,len(prices)
    for i := 0; i < n-1; i++ {
    	if prices[i] < prices[i + 1]{
    		res += prices[i+1] - prices[i]
    	}
    }
    return res
}