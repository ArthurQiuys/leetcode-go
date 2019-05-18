/*
* @Author: qiuyu
* @Date:   2019-05-18 23:35:51
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-18 23:48:30
*/
package main

func main() {
	
}
func convert(s string, numRows int) string {
    if len(s)==0 || numRows <= 0{
    	return s
    }
    if numRows == 1{
    	return s
    }
    res := []rune(s)
    size := 2* numRows-2
    for i := 0; i < numRows; i++ {
    	for j := i; j< len(s); j+= size{
    		res += s[j]
    		tmp := j+ size -2*i
    		if i !=0&& i!= numRows -1 && tmp < len(s) {
    			res += s[tmp]
    		}
    	}
    	return res
    }
    return res
}