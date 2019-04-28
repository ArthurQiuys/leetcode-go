/*
* @Author: qiuyu
* @Date:   2019-04-28 19:56:05
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-28 19:56:09
*/
func maxSubArray(nums []int) int {
    res := math.MinInt32
    sursum := 0
    for _ , i := range nums{
        sursum = max(sursum+i,i)
        res = max(res,sursum)
    }
    return res
}
func max( num1 int, num2 int)int{
    if num1> num2{
        return num1
    }
    return num2
}