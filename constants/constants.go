// 常数
package main

import "fmt"
const Pi = 3.14

func main(){
	const world ="世界"
	fmt.Println("hello",world)
	fmt.Println("happy",Pi,"day")

	const truth = true
	fmt.Println("right?",truth)
}