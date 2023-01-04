package main

import "fmt"

/*
求函数值
	已知 f(1)=3  f(n) = 2*f(n-1)+1
	请使用递归的思想编程，求出f(n)的值?
*/

func expression(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*expression(n-1) + 1
	}
}

func main() {
	res := expression(5)
	fmt.Println(res)
}
