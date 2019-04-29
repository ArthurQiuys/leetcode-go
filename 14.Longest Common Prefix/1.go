/*
* @Author: qiuyu
* @Date:   2019-04-26 03:19:45
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-30 00:15:59
*/
package main

func main() {
	
}
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }
    commonfix := strs[0]
    for i:= 0 ;i < len(strs); i++{
        if len(strs[i]) < len(commonfix){
            commonfix = strs[i]
            }
    }
    var prefix []byte
    for i:= range commonfix{
        for _, s:= range strs{
            if s[i] != commonfix[i]{
                return string(prefix)
            }
        }
        prefix = append(prefix, commonfix[i])
    }
    return string(prefix)
}
// 4 ms