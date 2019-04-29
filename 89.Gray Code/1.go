/*
* @Author: qiuyu
* @Date:   2019-04-29 19:57:07
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-29 23:14:37
*/
package main
import (
	"math"
)

func main() {
	i:= 2
	j:=grayCode(i)
	println(j)
}
func grayCode(n int) []int {
     res := make([]int, 0)
    for i := 0; i < int(math.Pow(float64(2), float64(n))); i++ {
    	// last :=((i>>1) ^ i)
    	// println(last)
    	res = append(res, (i>>1) ^ i)
    }
    return res
}