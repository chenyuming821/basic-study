// 1. 变量的声明
// package main

// import "fmt"

// var c, python, java bool

// // main 函数是程序的入口点
// func main() {
// 	// 在函数内部声明 int 类型变量 i，未初始化时默认为 0
// 	var i int
// 	// 打印变量值：i（局部变量）和 c、python、java（包级变量）
// 	// Go 中未初始化的变量会被赋予各自类型的零值
// 	fmt.Println(i, c, python, java)
// }

// 2. 变量的初始化
// package main

// import "fmt"

// var i,j int = 1,2

// // main 函数是程序的入口点
// func main() {
// 	// 在函数内部声明并初始化多个变量
// 	// Go 会根据右值自动推断变量类型
// 	var c, python, java = true, false, "no!"
// 	// 打印包级变量 i、j 和局部变量 c、python、java
// 	fmt.Println(i, j, c, python, java)
// }

// 3.简短变量声明
package main

import "fmt"

// main 函数是程序的入口点
func main() {
	// 使用 var 声明并初始化 int 类型变量 i 和 j,使用范围任何位置
	var i, j int = 1, 2
	// 使用 := 短声明语法声明并初始化变量 k（自动推断为 int 类型），仅函数内部有效
	k := 3
	// 使用 := 同时声明并初始化多个不同类型的变量
	c, python, java := true, false, "no!"

	// 打印所有变量的值
	fmt.Println(i, j, k, c, python, java)
}