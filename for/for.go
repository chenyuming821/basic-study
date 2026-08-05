// for 循环
package main

import "fmt"

// 方式一
// func main(){
// 	sum := 0
// 	for i:=0;i<10;i++{
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }


// for 循环的其他形式
func main(){
	sum := 1
	for ;sum<10;{//sum两边的分号也可以去掉
		sum +=sum
	}
	fmt.Println(sum)
}

//无线循环
// package main

// func main(){
// 	for{
// 	}
// }