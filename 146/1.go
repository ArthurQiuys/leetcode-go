/*
* @Author: qiuyu
* @Date:   2019-05-03 19:04:27
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-05-03 23:14:25
*/
package main

func main() {

}

type LRUCache struct {
    head *node
    tail *node
    store map[int] *node
    cap int
}

type node struct {
	prev *node
	next *node
	key int
	val int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{cap:capacity, store:make(map[int]*node)}
}

func (this *LRUCache) removeFromchain(n *node){//出栈
	 if n == this.head {
        this.head = n.next
    }
    
    if n == this.tail {
        this.tail = n.prev
    }
    
    if n.next != nil {
        n.next.prev = n.prev
    }
    
    if n.prev != nil {
        n.prev.next = n.next
    }
}

func (this *LRUCache)addTochain(n *node) {//入栈
	n.prev = nil
	n.next = this.head
	if n.next !=nil{
		n.next.prev = n
	}
	this.head = n
	if this.tail == nil{
		this.tail = n
	}
}
func (this *LRUCache) Get(key int) int {
    n ,ok := this.store[key]
    if !ok{
    	return -1
    }
    this.removeFromchain(n)
    this.addTochain(n)
    return n.val
}


func (this *LRUCache) Put(key int, value int)  {
    n, ok:= this.store[key]
    if !ok{
    	n = &node{val:value ,key : key}
    	this.store[key] = n
    }else{
    	n.val = value
    	this.removeFromchain(n)
    }
    this.addTochain(n)
    if len(this.store)> this.cap{
    	n = this.tail
    	if n!= nil{
    		this.removeFromchain(n)
    		delete(this.store, n.key)
    	}
    }
}
