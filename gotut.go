/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main

import(
	f "fmt"
	"bufio"
	"os"
	"log"
)

var pl = f.Println

//comment
/* 	block 
	comment */

func main(){
	pl("Hello go")
	f.Println("Hello go") //same as above
	//fmt.Print("Hello go") same as above but does not print in new line
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	//name, _ := reader.ReadString('\n') same as above but trivialy you cant handle errors this way
	if err == nil { // nil is returned when there is no error
		pl("Hello", name)
	}else {
		log.Fatal(err)
	}
}
