/*
* @Author: qiuyu
* @Date:   2019-04-24 02:01:32
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-24 02:01:59
*/
import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    var g , h , x float64
    if m < n {
        return findMedianSortedArrays(nums1 , nums2)
    }
    if m == 0{
        g = float64(nums1[(m-1)/2])
        h = float64(nums1[m/2])
        x = ( g + h )/ 2.0
        return x
    }
    left := 0
    right := n * 2
    var L1 , L2 , R1 , R2 float64
    for ; left < right ;{
        mid2 := ( left + right ) / 2
        mid1 := m + n - mid2
        if (mid1 == 0){
            L1 = -1
        } else {
            L1 = float64(nums1[(mid1 - 1) / 2])
        }
        if (mid2 == 0){
            L1 = -1
        } else {
            L1 = float64(nums2[(mid2 - 1) / 2])
        }
        if (mid1 == m * 2){
            R1 = math.MaxInt32
        } else {
            R1 = float64(nums1[mid1 / 2])
        }
        if (mid1 == n * 2){
            R2 = math.MaxInt32
        } else {
            R2 = float64(nums2[mid2 / 2])
        }
        if L1 > R2 {
            left = mid2 +1
        } else if L2 > R1{
            right = mid2 -1
        } else {
            return (math.Max( L1 , L2 ) + math.Min( R1 , R2 ) )/2
        }
    }
    return 0
}