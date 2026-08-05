// 命名返回值
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum / 9
	y = sum * 8
	return
}

func main(){
	fmt.Println(split(18))
}