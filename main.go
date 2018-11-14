package main

import "fmt"

func main()  {
	 mySlice  := [4]string{"this", "is", "my", "slice"}

	for index, value := range mySlice {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Println("\n++++++++++++++++++++++++++++\n")

	for index := range mySlice{
		fmt.Println("we can get only index", index)
	}

	fmt.Println("\n++++++++++++++++++++++++++++\n")

	for _, value := range mySlice{
		fmt.Println("or only value, with using _", value)
	}
}