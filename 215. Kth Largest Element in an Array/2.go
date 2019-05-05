/*
* @Author: qiuyu
* @Date:   2019-05-05 14:31:15
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:57:19
*/
package main

func main() {
	
}
func findKthLargest(nums []int, k int) int {
    h := &IntHeap{}
    heap.Init(h)
    for i := 0; i < k; i++ {
        heap.Push(h, nums[i])
    }
    for i := k; i < len(nums); i++ {
        if nums[i] < (*h)[0] {
            continue
        }
        heap.Pop(h)
        heap.Push(h, nums[i])
    }
    return (*h)[0]
}
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}