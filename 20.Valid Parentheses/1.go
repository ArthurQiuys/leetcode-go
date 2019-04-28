/*
* @Author: qiuyu
* @Date:   2019-04-27 02:13:27
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:18:11
*/
func isValid(s string) bool {
    stack := []rune{}
    m := map[rune]rune{
        ')':'(',
        ']':'[',
        '}':'{',
    }
    for _, i := range s{
        if len(stack) == 0{
            stack = append(stack , i)
        }else if i =='('|| i == '['|| i == '{'{
            stack = append(stack , i)
        }else{
            if m[i] != stack[len(stack) -1]{
                return false
            }else{
                stack = stack[:len(stack)-1]
            }
        }
    }
    return len(stack) == 0
}