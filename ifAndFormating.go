package main

import(
	"fmt"
	//"reflect"
	//"strconv"
	//"bufio"
	//"os"
	//"log"
)

func main(){
	iAge := 8
	if (iAge >= 1) && (iAge <= 18){
		fmt.Println("Important Birthday")
	} else if (iAge == 21) || (iAge == 50){
		fmt.Println("Also an important Birthday")
	} else if (iAge >= 65){
		fmt.Println("Important Birthday for retire")
	} else {
		fmt.Println("Not an important birthday")
	}

	/*
	-----------------------------------------Print Formats-----------------------------------------------------------
	%d Integer
	%c Character
	%f Float
	%t Boolean
	%s String
	%o Base 8
	%x Base 16
	%v Wildcard, it guesses based on datatype
	%T Type of supplied value
	*/

	fmt.Printf("%s %d %c %f %t %o %x\n", "Something", 16, 'A', 10.0/3.0, (2+2)==5, 1, 1)
	//the line below prints spaces on unused/extra decimal places
	fmt.Printf("Have a total of 9 decimals worth of precision %9f\n", 3.141592)
	fmt.Printf("Round the float to 2 decimal places %.2f\n", 10.0/3.0)
	stringFormattedPi := fmt.Sprintf("%.0f", 3.141592)
	fmt.Println(stringFormattedPi)
}