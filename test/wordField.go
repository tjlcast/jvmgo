package main

import (
	"fmt"
	"os"
	"log"
)

func f() {}
var g = "g"
var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main() {
	//f := "f"
	fmt.Println(f) // 如果是函数，则打印函数地址
	fmt.Println(g)
	//fmt.Println(h)

	x := "hello!"
	for i:=0; i<len(x); i++ {	// 隐式词法区 和 执行体词法域
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO"
		}
	}
}
