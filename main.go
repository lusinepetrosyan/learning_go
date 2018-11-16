package main

import "fmt"

func main() {
	var anyArray [4]string

	myArr := [5]int{1, 2, 3, 4, 5}
	myFunc(myArr, anyArray)


}

func myFunc(arr [5]int, anyArr [4]string) {
	fmt.Println("this is will print copy of array, Not reference of array", arr,"\n and this is empty array of string:", anyArr)
}
