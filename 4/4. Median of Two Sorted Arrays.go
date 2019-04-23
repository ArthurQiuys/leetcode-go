/*
* @Author: qiuyu
* @Date:   2019-04-23 23:40:58
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-24 02:08:54
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var i,j,mdx,pre,cur int
    l1 := len(nums1)
    l2 := len(nums2)
    lt := l1 + l2
    mdx = lt/2 + 1
    for (i+j) != mdx {
        pre = cur
        if i < l1 && j < l2 && nums1[i] <= nums2[j] {
            cur = nums1[i]
            i++
        } else if j < l2 {
            cur = nums2[j]
            j++
        } else if i < l1 {
            cur = nums1[i]
            i++
        }
    }
    if lt%2 == 1 {
        return float64(cur)
    }
    return float64(pre+cur)/2
}
//执行用时 : 28 ms, 在Median of Two Sorted Arrays的Go提交中击败了97.98% 的用户
//内存消耗 : 5.7 MB, 在Median of Two Sorted Arrays的Go提交中击败了64.36% 的用户