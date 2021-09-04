package main

import (
	"container/list"
	"errors"
	"fmt"
)

//最近没访问的淘汰

type lru struct{
	list  *list.List
	warehouse map[string]*list.Element
	capacity int
}

func NewLru(capacity int)*lru {
	return &lru{
		list : list.New(),
		warehouse: make(map[string]*list.Element),
		capacity: capacity,
	}
}

func (l *lru)Get(key string)(*list.Element, error) {
	 if value,ok := l.warehouse[key]; ok {
	 	 l.list.MoveToFront(value)
	 	 return  value, nil
	 }
	 return nil, errors.New("no value found")
}

func (l *lru)Set(key string, value interface{}) {
	    l.list.PushFront(value)
		l.warehouse[key] = l.list.Front()
		//todo 删除过期元素
		for l.list.Len() > l.capacity {
			l.list.Remove(l.list.Back())
		}
}

func (l *lru)travel(){
	for i := l.list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	return
}

func main() {
	  var l = NewLru(5)
	  l.Set("25",12)
	  l.Set("12",24)
	  l.Set("test",8)
	  l.Set("yes",1)
	  l.Set("test121",7)
	  l.Get("25")
	  l.Set("no",30)
	  l.travel()

}