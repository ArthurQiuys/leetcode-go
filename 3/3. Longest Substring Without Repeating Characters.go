/*
* @Author: qiuyu
* @Date:   2019-04-23 23:31:57
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-23 23:32:47
*/
func lengthOfLongestSubstring(s string) int {
    array := [256]int{}
	for i := 0; i < len(array) ; i++ {
		array[i] = -1
	}
    length := 0
    pre := -1
    cur := 0
    
    for i := 0; i < len(s) ; i++ {
        pre = max(pre, array[s[i]])
        cur = i - pre
        length = max(length, cur)
        array[s[i]] = i
    }
    
    return length
}

func max(a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
    
}
//Runtime: 8 ms, faster than 84.89% of Go online submissions for Longest Substring Without Repeating Characters.
//Memory Usage: 2.6 MB, less than 88.34% of Go online submissions for Longest Substring Without Repeating Characters.