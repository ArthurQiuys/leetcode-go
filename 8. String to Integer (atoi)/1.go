/*
* @Author: qiuyu
* @Date:   2019-04-25 02:03:18
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-25 02:03:32
*/
func myAtoi(str string) int {
    MAX, MIN := (1<<31)-1, 1<<31
    if len(str) == 0{
        return 0
    }
    slice := []rune(str)
/*    for char := range str{
        slice = append( slice , rune(char))
    }*/
    sign ,base ,i , n:= 1,0,0,len(str)
    for i<n && slice[i] == ' '{
        i++
    }
    if i < n && ( slice[i] == '+'|| slice[i] == '-'){
        if slice[i] == '+'{
            sign = 1
        }else{
            sign = -1
        }
        i++
    }
    for  i < n && slice[i] >='0' && slice[i] <= '9'{
        t := int (slice[i] - '0')
        if sign >0 && base > (MAX-t)/10|| sign < 0 && base>((MIN-t)/10){
            if sign ==1{
                return MAX
            }else{
                return -MIN
            }
        }
        base = 10 * base + t
        i++
    }
    return base * sign
}