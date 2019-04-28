/*
* @Author: qiuyu
* @Date:   2019-04-27 02:12:13
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:12
*/
func threeSum(nums []int) [][]int {
    lens := len(nums)
    ans := make([][]int,0)
    if lens < 3 {
        return ans
    }
    sort.Ints(nums)
    for k, _:= range nums{
        if k > 0 && nums[k] == nums[k-1]{
            continue
        }
        l := k + 1
        r := lens -1
        for l < r{
            sum := nums[k] + nums[l] + nums[r]
            if sum == 0{
                array := []int {nums[k], nums[l], nums[r]}
                ans = append (ans ,array)
                l ++
                r --
                for l < r &&nums[l] == nums[l - 1]{
                    l++
                }
                for l < r && nums[r] == nums[r + 1]{
                    r--
                }
            }else if sum >0{
                r--
            }else if sum < 0{
                l++
            }
        }
    }
    return ans
}