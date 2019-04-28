/*
* @Author: qiuyu
* @Date:   2019-04-26 03:21:29
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:12
*/
func maxArea(height []int) int {
    res , i , j := 0 , 0 , len(height)-1
    for i < j{
        h := int(math.Min(float64(height[i]),float64(height[j])))
        res = int(math.Max(float64(res),float64(h*(j - i))))
                  for i < j && h == height[i]{
                      i++
                  }
                  for i < j && h == height[j]{
                      j--
                  }
    }
                  return res
}
//24 ms