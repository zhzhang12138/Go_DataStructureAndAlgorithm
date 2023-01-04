package main

import "fmt"

/*
斐波那契数列是一种数列，其中的每一项都是前两项的和。斐波那契数列的前几项如下：
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...

给你一个整数n，求出它的斐波那契数是多少？
	思路：
		(1)当n==1||n==2，返回1
		(2)当n>2，返回前面两个数的和 f(n-1)+f(n-2)
*/

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	res := fibonacci(6)
	fmt.Println(res)
}
