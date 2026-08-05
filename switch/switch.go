package main

import (
	"fmt"
	//"runtime"
	"time"
)

//1。判断操作系统
/*/
func main() {
	fmt.Print("Go runs on ")
	switch os :=runtime.GOOS;os{//runtime.GOOS来判断操作系统
		case "darwin":
			fmt.Println("macOS")
		case "linux":
			fmt.Println("Linux")
		default://default3分支确定Windows系统
			fmt.Printf("%s.\n",os)
	}
}
/*/

// 2. 判断当前时间距离周六有多久
func main() {
	fmt.Println("When's Saturday?")
	// 获取当前日期是星期几（返回值类型为 time.Weekday）
	today := time.Now().Weekday()
	// switch 表达式是 time.Saturday（常量，值为6）
	// 每个 case 会与 today+偏移量 进行比较
	switch time.Saturday {
		case today + 0:  // 如果 Saturday == today，说明今天就是周六
			fmt.Println("Today is Saturday.")
		case today + 1:  // 如果 Saturday == today+1，说明明天是周六
			fmt.Println("Tomorrow is Saturday.")
		case today + 2:  // 如果 Saturday == today+2，说明后天是周六
			fmt.Println("In two days, it will be Saturday.")
		default:  // 其他情况，距离周六超过两天
			fmt.Println("We're not there yet.")
	}
}