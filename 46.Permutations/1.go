/*
* @Author: qiuyu
* @Date:   2019-04-28 19:55:06
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-28 19:55:17
*/
func permute(nums []int) [][]int {
    res := make([][]int ,0)
    if len(nums) == 1{
        res =append(res,nums)
        return res
    }
    for i := range nums{
        subnum := make([]int, len(nums))
        copy(subnum, nums)
        subnum = append(subnum[:i], subnum[i+1:]...)
        subR:= permute(subnum)
        for _, sr := range subR{
            rr := []int{nums[i]}
            rr = append(rr,sr...)
            res = append(res,rr)
        }
    }
    return res
}