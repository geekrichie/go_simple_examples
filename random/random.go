package main

import (
	"crypto/rand"
	"fmt"
	"unsafe"
)

func randNew() {
	//r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Printf("%v\n", r.Intn(100))
	//fmt.Printf("%v\n", r.Intn(100))
	//fmt.Printf("%v\n", r.Intn(100))
	//fmt.Printf("%v\n", r.Intn(100))
	//fmt.Printf("%v\n", r.Intn(100))
}

func randTrue() {
	b := make([]byte, 4)
	n, err := rand.Read(b)
	fmt.Println(n)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	//824633802936
	fmt.Printf("%d\n",*(*int)(unsafe.Pointer(&b[0])) )
}

func main() {
	//randTrue()
	//var p = make([]int,4)
	//fmt.Printf("%p\n",p)
	//fmt.Printf("%p\n",&p)
	//fmt.Printf("%p\n",&p[0])
	//sh := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	//fmt.Printf("A1 Data:%x,Len:%d,Cap:%d\n",sh.Data,sh.Len,sh.Cap)
	type T struct {
		msg string
	}

	var g *T
	if g == nil {
		println("hello")
	}
}