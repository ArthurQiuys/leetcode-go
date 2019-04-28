/*
* @Author: qiuyu
* @Date:   2019-04-27 02:15:33
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:04
*/
func search(nums []int, target int) int {
    left , right := 0, len(nums)-1
    for ;left <= right;{
        mid := left + (right - left)/2
        if nums[mid] == target{
            return mid
        }
        if nums[mid] > nums[right]{
            if nums[mid] > target && nums[left]<= target{
                right = mid -1
            }else{
                left = mid +1
            }
        }else{
            if nums[mid] < target && nums[right] >= target{
                left = mid +1
            }else{
                right = mid-1
            }
        }
    }
    return -1
}