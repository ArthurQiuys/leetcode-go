/*
* @Author: qiuyu
* @Date:   2019-04-29 23:24:58
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 23:42:55
*/
package main
import (
	"math"
)
func main() {
	t:= []int {7,5,3,6,4}
	a:=maxProfit(t)
	println("最后的☞",a)
}
func maxProfit(prices []int) int {
    res ,buy:=0, math.MaxInt32
    for _,i := range prices{
    	buy = min(buy , i)
    	res = max(res, i - buy)
    }
    return res
}
func min(a int, b int)int {
	if a > b {
		println("最小值为",b)
		return b
	}
	println("最小值为",a)
	return a
}
func max(a int , b int)int {
	if a> b{
		println("最大值为",a)
		return a
	}
	println("最大值为",b)
	return b
}