/*
* @Author: qiuyu
* @Date:   2019-04-25 02:02:45
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-25 02:03:07
*/
import"math"
func reverse(x int) int {
    res := 0
    for  ; x!=0 ;{
        if math.Abs(float64(res))>(math.MaxInt32/10){
            return 0;
        }
        res = res * 10 + x % 10;
        x /= 10
    }
    return res
}