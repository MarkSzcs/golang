package main

import(
	"fmt"
	"reflect"
	"strconv"
	//"bufio"
	//"os"
	//"log"
)

// importing something and then not
// using it -> compilation fail

var p1 = fmt.Println

func main(){
	// variable declaration syntax
	// var variable_name type
	/* sidenote: pay attention
	   starting your variable name
	   with capital letters
	   might result in unwanted
	   sideeffects
	   (that var can be accessed
		even outside the package)
	*/
	var firstName string = "Mark"
	p1(firstName)
	p1(reflect.TypeOf(firstName))
	// what if i dont declare a type?
	// nothing, it is completely fine
	var lastName = "Szucs"
	p1(lastName)
	p1(reflect.TypeOf(firstName))
	//in both cases the typeof
	//returns a string


	//shorthand NEW variable declaration
	//middleName := ""
	//middleName = "Szilard"

	// variable of a given type have
	// default values:
	// string = ""
	// int = 0
	// bool = false
	// float = 0.0
	// rune = 0

	var myRune rune
	p1(myRune)
	
	//typecast is as simple
	//as the following in 
	//most cases
	p1(int(myRune))

	//string to int is a lil different
	conv := "5000"
	//it requires some errorhandling
	toInt, err := strconv.Atoi(conv)
	p1(toInt,err,reflect.TypeOf(toInt))

	//int to string
	//you need no errorhandling
	toStr := strconv.Itoa(toInt)
	p1(toStr, reflect.TypeOf(toStr))

	//string to float
	fStr := "3.14"
	//some advanced errorhandling
	//when defining multiple variables
	//and using := operator
	//only one needs to be a new var
	if fStr,err := strconv.ParseFloat(
	   fStr, 64); err == nil{
	   p1(fStr)
	}

	//float to string, crooked way
	fToString := fmt.Sprintf("%f", 3.14)
	p1(fToString)
	//first if statement in go
	//the only correct syntax is
	//{ being immediately after the
	//if statement else -> error

}