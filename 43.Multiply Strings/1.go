/*
* @Author: qiuyu
* @Date:   2019-04-27 02:15:46
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-27 02:16:12
*/
func multiply(num1 string, num2 string) string {
    res := make([]int,len(num1)+len(num2))
    for i:= len(num1) -1; i>=0 ;i--{
        for j:= len(num2)-1; j>=0; j--{
            a := int(num1[i] - '0')
            b := int(num2[j] - '0')
            c := a*b +res[i+j+1]
            res[i+j] += c/10
            res[i+j+1] = c % 10
        }
    }
    k :=0
    l := 0
    for ; l<len(res); l++{
        if res[l] !=0 {
            k = l
            break
        }
    }
    if l == len(res){
        return "0"
    }
    var s strings.Builder
    for i:= k ;i<len(res);i++{
        s.WriteString(string(res[i] + '0'))
    }
    return s.String()
}