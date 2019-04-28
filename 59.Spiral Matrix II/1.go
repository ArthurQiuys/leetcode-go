/*
* @Author: qiuyu
* @Date:   2019-04-28 21:01:22
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-28 21:26:05
*/
package main

func main() {
	
}
func generateMatrix(n int) [][]int {
    res := make([][]int, n)
    for i:= range res {
    	res[i] = make([]int, n)
    }
    if n <1 {
    	return res
    }

    value , up, down, left ,right := 1 , 0 , n-1 ,0 ,n-1
    for up<= down && left <= right {
    	for i := left; i <= right; i++ {
    		res[up][i]= value 
    		value ++
    	}
    	up ++
    	for i := up; i <= down; i++ {
    		res[i][right]= value
    		value ++
    	}
    	right--
    	if up <= down{
    		for i := right; i >= left; i-- {
    			res[down][i]= value 
    			value ++
    		}
    		down --
    	}
    	if left <= right{
    		for i := down; i >= up; i-- {
    			res[i][left]=value 
    			value ++
    		}
    	}
    	left ++
    }
    return res
}