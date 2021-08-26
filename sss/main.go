package main

import "User/reptile"

//import (
//	"fmt"
//)
//
//func do(i interface{}){
//	switch v:= i.(type) {
//	case int:
//		fmt.Println("我是数字",v)
//	case string:
//		fmt.Println("我是字符串",v,len(v))
//	default:
//		fmt.Println("我是其他",v)
//	}
//}
//
//func main() {
//	do("2")
//	do(12)
//	do(true)
//}

//type I interface {
//	M()
//}
//
//
//type T struct {
//	S string
//}
//
//func (t T) M() {
//	fmt.Println(t.S)
//}
//
//type F interface {
//	a()
//}
//
//type S struct {
//	k string
//}
//
//func (s1 S) a() {
//	fmt.Println(s1.k)
//}
//
//func main() {
//	var i I = T{"hello"}
//	i.M()
//
//	var f F = S{"世界"}
//	f.a()
//}



func main() {
	reptile.Data_get()
}
















