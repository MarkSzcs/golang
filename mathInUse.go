package main

import(
	"fmt"
	"bufio"
	//"math"
	//"strings"
	"strconv"
	
)

var pl = fmt.Println

func main(){

	reader := bufio.NewReader(os.Stdin)
	// first num
	s_n1, err1 := reader.ReadString('\n')
	// second num
	s_n2, err2 := reader.ReadString('\n')
	
	// output its sum
	if(err1 == nil || err2 == nil){
		s_n1 = strings.TrimSpace(s_n1)
		s_n2 = strings.TrimSpace(s_n2)
		n1, err1 := strconv.Atoi(s_n1)
		n2, err2 := strconv.Atoi(s_n2)
		if(err1 == nil || err2 == nil){
			pl("The sum of the input numbers is: ", n1+n2)
		}
		else {
			log.Fatal(err1 + err2)
			os.Exit(-3)
		}
	}
	else {
		log.Fatal(err1 + err2)
		os.Exit(-2)
	}
}	