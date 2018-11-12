package main

import "fmt"

func main()  {
	var a string
	var (  //other type of declaration variable
		b int
	 	c bool
	)
	d:= 1.4 //short declaration type of go

	a, b, c = "apple", 12, true

	fmt.Println("in printf %v can contains any existing type", a, b, c, d)
	fmt.Println("%T type of variable i printf \n")
	fmt.Println("%g float number but removes unnecessary fraction part \n")
	fmt.Println("println no need to new line, it automatically sets")


	fmt.Printf("my string is: %s \n my bool is: %t \n my int is: %d \n my float is: %f \n\n ", a, c, b, d)


	fmt.Printf("float number %f \n", d)
	fmt.Printf("float number without unnecessary fraction part %g \n\n", d)



	fmt.Println("new line character is \\n")
	fmt.Println("double cotes is \" haha done it \" ")


}