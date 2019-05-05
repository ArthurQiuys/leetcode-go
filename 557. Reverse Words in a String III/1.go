/*
* @Author: qiuyu
* @Date:   2019-05-05 18:46:32
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-05 18:53:02
*/
package main

func main() {
	

}
func reverseWords(s string) string {
    ls := len(s)
    if ls < 2{
    	return s
    }
    b := []byte(s)
    l,r:=0,0
    for i , v:= range b{
    	if v ==' '|| i ==ls -1{
    		r = i - 1
    		if i == ls -1{
    			r ++
    		}
    		for l< r{
    			b[l] , b[r] = b[r] , b[l]
    			r--
    			l++
    		}
    		l = i + 1
    	}
    }
    return string(b)
}