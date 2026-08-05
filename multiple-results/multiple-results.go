// 多返回值
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y,x
}//换成其他的类型也可以，比如int, float64, bool等,之所以要写两个string，是因为有两个返回值

// main 函数是程序的入口点
func main() {
	// 调用 swap 函数交换两个字符串，将返回值分别赋值给变量 a 和 b
	a, b := swap("hello", "world")
	// 打印交换后的结果
	fmt.Println(a, b)
}