package basics

import "fmt"

func Pointer() {
	//Go语言支持指针
	//不过不支持指针运算
	a, b := 3, 5
	pa, pb := &a, &b
	fmt.Println(*pa, *pb)
}
