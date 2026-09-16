	
package main

import "fmt"

/*func main(){
	defer fmt.Println("hello world")
	fmt.Println("hello go")

}*/

// defer multi
func main(){
	fmt.Println("start")
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
fmt.Println("end")
	
}