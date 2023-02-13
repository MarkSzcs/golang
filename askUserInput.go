package main

import (
	"fmt"
	"strings"
	b "bufio"
	"os"
	"strconv"
)

var pl = fmt.Println

func main(){
	//accept a string as string as an input like: "Name Age"
	//format the string
	//decide what kind of education should the holder of the Name with Age recieve

	correctInput := false
	for correctInput != true {
		reader := b.NewReader(os.Stdin)
		strInput, err1 := reader.ReadString('\n')
		if (err1 == nil){
			strAsList := strings.Split(strInput, " ")
			name := strAsList[0]
			sAge := strAsList[1]
			age, err2 := strconv.ParseInt(sAge)
			if (err2 == nil){
				if (age < 5){
					pl(name, " You are young enjoy it!")
				}
				switch age {
					case 5:
						pl(name, " Go to kindergarten!")
					case 6:
						pl(name, " You belong to the 1st class of elementary school.")
					case 7:
						pl(name, " You belong to the 2nd class of elementary school.")
					case 8:
						pl(name, " You belong to the 3rd class of elementary school.")
					case 9:
						pl(name, " You belong to the 4th class of elementary school.")
					case 10:
						pl(name, " You belong to the 5th class of elementary school.")
					case 11:
						pl(name, " You belong to the 6th class of elementary school.")
					case 12:
						pl(name, " You belong to the 7th class of elementary school.")
					case 13:
						pl(name, " You belong to the 8th class of elementary school.")
					case 14:
						pl(name, " You belong to the 9th class of high school.")
					case 15:
						pl(name, " You belong to the 10th class of high school.")
					case 16:
						pl(name, " You belong to the 11th class of high school.")
					case 17:
						pl(name, " You belong to the 12th class of high school.")
					default:
						pl(name, " You should go to college!")
				}
			}
		}
	}
}