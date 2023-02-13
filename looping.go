package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	//format
	// for initialization; contidion; statement {BODY}
	for x :=1; x<=5; x++ {
		pl(x)
	}

	for x :=5; x>=1; x-- {
		pl(x)
	}

	//basic stuff
	// the variable x trivially does not exist outside of the scope
	
	//while loop not so trivially is also simulated by a for loop, there is no while keyword
	fX := 1
	for fX <= 5{
		pl(fX)
		fX++;
	}

	//array traversal
	aNums := []int{1,2,3}
	for _, num := range aNums {
		pl(num)
	}

	//do while also simulated by a for loop, there is no do-while keyword
	ind := 0
	for true{
		if (aNums[ind] % 2) == 0{
			pl("The firs even number in the array is:",aNums[ind],"at the index of:",ind)
			//pl("At the index of: ", ind)
			break
		} else {
			if len(aNums) <= ind {
				pl("There are no even numbers in the array.")
				break
			}
		}
		ind++
	}


}