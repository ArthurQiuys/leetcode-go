/*
* @Author: qiuyu
* @Date:   2019-04-27 02:15:17
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:17:57
*/
func removeDuplicates(nums []int) int {
    if len(nums) == 0{
        return len(nums)
    }
    i , j ,n:= 0,0,len(nums)
    for j<n{
        if nums[i] == nums[j]{
            j++
        }else{
            i++
            nums[i] = nums[j]
            j++
        }
    }
    return i+1
}