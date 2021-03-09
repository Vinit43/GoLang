package main

import "fmt"


func Calculate (no1 int , no2 int) (int , int){

	var result1 = 0 
	var result2 = 0

	result1 = no1 + no2
	result2 = no2 - no1

	return result1 , result2      // GoLang supports MULTIPLE RETURN VALUE
} 

func main() {

	var value1 = 10
	var value2 = 20
	var ret1 = 0
	var ret2 = 0

	ret1,ret2 = Calculate(value1,value2)
	fmt.Println("Addition is:",ret1)
	fmt.Println("Substraction is:",ret2)

	
}