package main

import "fmt"

func main()  {
	myArr := [6]int{1,2,3,14,5,6}
	var t []int = myArr[1 : 4]  // 4 th index not included

	fmt.Println(t)
}