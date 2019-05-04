/*
* @Author: qiuyu
* @Date:   2019-05-04 18:13:19
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-04 18:23:39
*/
package main

import (
	"math"
)
func main() {
	
}
type MinStack struct {
    Val int
    Min int
    Vaild bool
    Next *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack(0 , 0 , false , nil)    
}


func (this *MinStack) Push(x int)  {
    if this.Vaild{
    	min := int(math.Min(float64(x), float64(this.Min)))
    	m := this
    	this.Min = min
    	this.Val = x
    	this.Vaild = true
    	this.Next = m
    } else{
    	this.Min = x
    	this.Val = x
    	this.Vaild = true
    }
}


func (this *MinStack) Pop()  {
    if this.Vaild{
    	if this.Next!=nil{
    		this = this.Next
    	}else{
    		this.Val = 0
    		this.Vaild = false
    		this.Min = 0
    	}
    }
}


func (this *MinStack) Top() int {
    return this.Val
}


func (this *MinStack) GetMin() int {
    return this.Min
}