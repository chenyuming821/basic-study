package main

import (
	"fmt"
	"math"
)

func pow(x,n,lim float64) float64{
	if v:=math.Pow(x,n);v<lim{
		return v
	}else{//else不可新立一行，否则会报错
		fmt.Printf("%g>=%g\n",v,lim)
	}
	return lim
}

func main(){
	fmt.Println(pow(2,3,10))
}