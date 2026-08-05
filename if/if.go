package main

import (
	"fmt"
	"math"
)
//if判断语句，根据判断条件执行不同的分支，if语句可以没有else分支，也可以有多个else if分支
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// func main() {
// 	fmt.Println(sqrt(2), sqrt(-4))
// }

//if语句用简短的陈述
// pow 函数计算 x 的 n 次方，并与上限值 lim 比较
// 参数:
//   x   - 底数（base）
//   n   - 指数（exponent）
//   lim - 上限值（limit）
// 返回值:
//   如果 x^n < lim，返回 x^n 的结果；否则返回 lim
func pow(x, n, lim float64) float64 {
	// Go 特有语法：在 if 条件前声明临时变量 v
	// v 的作用域仅限于 if/else 块内
	if v := math.Pow(x, n); v < lim {
		// 如果计算结果小于上限，返回计算值
		return v
	}
	// 如果计算结果大于等于上限，返回上限值
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}