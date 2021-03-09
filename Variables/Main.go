package main

import "fmt"

func main() {

	/*	
		GOLang gives error if you initialize a varibale and dont use it.

		So its a rule to use that initialized variable

	*/
	
	var no1 int = 11                // format for initializing variable is   var variable_name variable_type = data

	var no2 = 21

	no3 := 31

	var no4 int                  // If you dont declare the Data_Type of variable it gives error
	no4 = 41
   
 	const no5 = 51             // This is the constant variable which is not going to change through out the code

 	var no6 float32 = 11.3

 	var str string = "Vinit"

 	str1 := "Nandurkar"

	fmt.Println("First number is",no1)

	fmt.Println("Second number is",no2)

	fmt.Println("Third number is",no3)

	fmt.Println("Fourth number is",no4)

	fmt.Println("Fifth number is",no5)

	fmt.Println("Sixth number is",no6)

	fmt.Println("String name is",str)

	fmt.Println("String end is",str1)

	
}