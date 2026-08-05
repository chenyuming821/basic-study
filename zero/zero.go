// 零值初始化
package main

import "fmt"

// main 函数演示 Go 的零值初始化特性
func main() {
	// 声明 int 类型变量，未初始化时默认为 0
	var i int
	// 声明 float64 类型变量，未初始化时默认为 0.0
	var f float64
	// 声明 bool 类型变量，未初始化时默认为 false
	var b bool
	// 声明 string 类型变量，未初始化时默认为空字符串 ""
	var s string

	// 使用格式化输出打印各类型的零值
	// %v: 按默认格式输出值，%q: 输出带引号的字符串
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}