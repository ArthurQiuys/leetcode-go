/*
* @Author: qiuyu
* @Date:   2019-04-26 03:21:59
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:03
*/
func maxArea(height []int) int {
    res , i , j:= 0, 0, len(height)-1
    for i < j{
        res = int(math.Max(float64(res), math.Min(float64(height[i]), float64(height[j]))*float64(j - i)))
        if height[i] < height[j]{
            i++
        }else{
            j--
        }
    }
    return res
}
// 24ms