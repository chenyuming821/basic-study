package main

import "fmt"

/*func main() {
	prime := [6]int{2, 3, 5, 7, 11, 13}

	var s []int=prime[1:3]
	fmt.Println(s)
}*/

/*//slices-pointer
func main() {
	name := [4]string{"a","b","c","d"}

	fmt.Println(name)

	a := name[0:2]
	b := name[1:3]
	fmt.Println(a,b)

	b[0] = "e"
	fmt.Println(a,b)
	fmt.Println(name)
}*/

/*//slice-literals
//首先一组数组列举出六个数字，然后另一个数组列出六个bool值，分别打印两个数组，最后构建结构，打印结构
func main() {
	a := []int{1,2,3,4,5,6}
	fmt.Println(a)
	b := []bool{true,true,false,true,false,true}
	fmt.Println(b)

	c := []struct {
		 a int
		 b bool}
		 {
		{1,true},
		{2,true},
		{3,false},
		{4,true},
		{5,false},
		{6,true},
	}
	fmt.Println(c)
}*/

/*//slice-len-cap
//首先对已知数组进行切片，切成0长度，接着扩展切片长度，然后丢掉前两个元素进行切片，最后输出len和cap
func main() {
	a := []int{1,2,3,4,5,6}
	fmt.Println(a)

	b := a[0:0]
	fmt.Println(b)

	c := a[:4]
	fmt.Println(c)

	d := a[2:]
	fmt.Println(d)
	
	fmt.Printf("len(d) = %d, cap(d) = %d, d = %v\n", len(d), cap(d), d)
}*/

/*//nil-slice 零切片
//定义一个零数组，打印数组长度以及容量，并进行if判断，如果是nil则输出
func main() {
	var s []int
	fmt.Println(s,len(s),cap(s))

	if s == nil {
		fmt.Println("s is nil")
	}
}*/

//making-slice 切片的创建
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

