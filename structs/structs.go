package main

import "fmt"

type a struct {
	i, j int
}

/*func main(){
	fmt.Println(a{1,2})
}*/

// 结构体字段的访问与修改
/*func main() {
	s := a{1, 2}   // 创建结构体 a 的实例：i=1, j=2
	s.j = 3        // 通过点号访问并修改字段 i 的值
	fmt.Println(s.j) // 输出字段 i 的值：3
}*/

//struct pointer
/*func main() {
	s := a{1, 2}
	p := &s // 创建结构体 a 的指针
	p.j = 3 // 通过指针访问并修改字段 j 的值
	fmt.Println(p.j) // 输出字段 j 的值：3
}*/

//struct-literals
var (
	a1 = a{1, 2} // 创建结构体 a 的实例：i=1, j=2
	a2 = a{i: 1}
	a3 = a{}
	p = &a{i: 1, j: 2}
)
func main() {
	fmt.Println(a1, a2, a3, p)
}