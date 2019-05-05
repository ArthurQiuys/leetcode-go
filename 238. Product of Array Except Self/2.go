/*
* @Author: qiuyu
* @Date:   2019-05-05 18:10:10
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:57:16
*/
func productExceptSelf(nums []int) []int {
	total := 1
	zero := 0
	// sum
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			total *= nums[i]
		} else {
			// count 0
			zero += 1
		}
	}
	// divide
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// judge 0
			if zero == 0 {
				nums[i] = total / nums[i]
			} else {
				nums[i] = 0
			}
		} else {
			// judge count
			if zero > 1 {
				nums[i] = 0
			} else {
				nums[i] = total
			}
		}
	}
	return nums
}
