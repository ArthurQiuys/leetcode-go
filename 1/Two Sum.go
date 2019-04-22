func twoSum(nums []int, target int) []int {
        sum := make(map[int]int)
    for i, n := range nums{
        if val, ok:= sum[n]; ok{
            return []int{val, i}
        }
        sum[target - n] = i
    }
    return nil//error
}
//执行用时 : 12 ms, 在Two Sum的Go提交中击败了69.61% 的用户
//内存消耗 : 3.8 MB, 在Two Sum的Go提交中击败了40.73% 的用户