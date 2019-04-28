/*
* @Author: qiuyu
* @Date:   2019-04-27 02:12:51
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:11
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	sum, dataS := 0, math.MaxInt64
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < target:
				l++
				if dataS > target-s {
					dataS = target - s
					sum = s
				}
			case s > target:
				r--
				if dataS > s-target {
					dataS = s - target
					sum = s
				}
			default:
				return s
			}
		}
	}
	return sum
}