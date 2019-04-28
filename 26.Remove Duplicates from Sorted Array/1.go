/*
* @Author: qiuyu
* @Date:   2019-04-27 02:15:05
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:09
*/
func removeDuplicates(nums []int) int {
    if len(nums) == 0{
        return len(nums)
    }
    i := 1
    for j:=i;j<len(nums);j++{
        if nums[i-1] != nums[j]{
            nums[i] = nums[j]
            i++
        }
    }
    return i
}