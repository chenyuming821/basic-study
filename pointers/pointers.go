package main

import "fmt"

func main() {
	i,j:= 1,2
	p := &i
	fmt.Println(*p)

	*p = 2
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)

}