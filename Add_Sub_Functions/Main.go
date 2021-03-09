package main

import "fmt"

func Addition (no1 int , no2 int) int{

	var Ans int = 0

	Ans = no1 + no2

	return Ans
}

func Substraction (no1 int , no2 int) int{

	var Ans int = 0

	Ans = no1 - no2

	return Ans
}

func main() {

	var value1 = 10
	var value2 = 20
	var result = 0

	result = Addition(value1 , value2)
	fmt.Println("The Addition is:",result)


	result = Substraction(value1 , value2)
	fmt.Println("The Substraction is:",result)
}