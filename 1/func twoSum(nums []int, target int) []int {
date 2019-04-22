func twoSum(nums []int, target int) []int {
    var n []int
    num := len(nums)
    m := make(map[int]int)
    for i:=0; i<num ;i++{
        m[nums[i]]=i
    }
    for i:=0 ; i<num ;i++{
        if thenum,ok := m[target-nums[i]];ok{
            if thenum ==i{
                continue
            }
            n = append( n, i, thenum)
            return n
        }
    }
    return n
}
//执行用时 : 8 ms, 在Two Sum的Go提交中击败了97.03% 的用户
//内存消耗 : 3.8 MB, 在Two Sum的Go提交中击败了23.21% 的用户