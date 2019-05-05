/*
* @Author: qiuyu
* @Date:   2019-05-04 21:25:05
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 14:31:04
*/
package main

func main() {
	
}
func findKthLargest(nums []int, k int) int {
    left , right := 0, len(nums) -1
    for len(nums) !=0{
    	pos := partition(nums, left , right)
    	if pos ==k-1{
    		return nums[pos]
    	}else if pos >k-1{
    		right = pos-1
    	}else {
    		left = left +1
    	}
    }
    return 0
}

func partition( nums []int, left int , right int)int {
	pivot , l , r:= nums[left], left +1, right
	for l <= r {
		if nums[l] < pivot && nums[r] > pivot{
			swap(l+1, r-1)
			l++
			r--
		}
		if nums[l] >= pivot{
			l++
		}
		if nums[r] <= pivot{
			r--
		}
	}
	swap(left , r)
	return r
}
func (nums *[]int)swap(a int, b int){
	var temp int
	temp = nums[a]
	nums[a] = nums[b] 
	nums[b] = temp
}