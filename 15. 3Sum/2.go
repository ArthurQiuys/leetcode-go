/*
* @Author: qiuyu
* @Date:   2019-04-27 02:12:18
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:11
*/
func threeSum(nums []int) [][]int {
    var ans [][]int
    if len(nums) <=2{
        return ans
    }
    sort.Ints(nums)
    for i :=0; i+2 <len(nums); i++{
        if i>0 && nums[i] == nums[i -1]{
            continue
        }
        if nums[i] > 0 {
			break
		}
        left ,right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0{
                ans = append(ans, append([]int{},nums[i],nums[left],nums[right]))
                left++
                right--
                for left < right && nums[left]== nums[left -1]{
                    left++
                }
                for left < right && nums[right] == nums[right +1]{
                    right --
                }
            }else if sum<0 {
                left ++
            }else{
                right --
            }
        }
    }
    return ans
}