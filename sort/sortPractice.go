package main

import "fmt"

type StringSlice []string
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func main() {
	//var b = []string{"d", "b", "c"}
	//sort.Sort(StringSlice(b))
	//fmt.Println(b)
	//var maxUint uint32
	//var maxInt int
	//maxUint = ^uint32(0)
	//maxInt  = int(maxUint >> 1)
	//fmt.Printf("%d\n %d", maxUint, maxInt)
	var  p = make([]int, 2)
    p[0] = 1
    var x = &p
    var y = make([]int,len(*x))
    copy(y, p)
    y[0] =2
	fmt.Println(p)
    fmt.Println(y)
}
