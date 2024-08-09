package main

import "fmt" //package  name

func main() {

	fmt.Print("hello, ") //simple function to print
	fmt.Print("world! \n")

	fmt.Println("  ")
	fmt.Println("Hello Sikan-dar") //simple function to print line

	//using variables
	var firstVar string = "Artreus"
	var secVar = "naruto"
	var thirdVar string

	fmt.Println(firstVar, secVar, thirdVar)

	firstVar = "kratos"
	thirdVar = "sakura"

	fmt.Println(firstVar, secVar, thirdVar)

	fourVar := 123
	fiveVar := "Ghost" //can not use this method of decalaring outside of func

	fmt.Println(fourVar, fiveVar)

	//ints var

	var num1 int = 10
	var num2 = 20
	num3 := 30

	var num5 float32 = 26.98
	num4 := 25.98 //this will pick float64 by default

	fmt.Println(num1, num2, num3, num4, num5)

	//using vars with text

	age := 26
	name := "sam"

	fmt.Println("Hello, My name is", name, "and I am", age, "years old.")
	fmt.Printf("Hello, My name is %v and I am %v years old. \n", name, age) //formatted strings
	// %_  = format specifier

	fmt.Printf("Hello, My name is %q and I am %v years old. \n", name, age)
	fmt.Printf("you scored %f points and he scored %0.2f points..!", 29.99, 29.99)
}
