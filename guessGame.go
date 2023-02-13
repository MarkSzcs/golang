package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math/rand"
	"time"
)

var pl = fmt.Println

//func tellMe(theIntVal int) int{
//	return theIntVal
//}

func main(){
	pl("Guess an integer between 1 and 100, at least try once ;)")
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	randNum := rand.Intn(100)+1

	reader := bufio.NewReader(os.Stdin)
	userGuess,err := reader.ReadString('\n')
	if err != nil {
		os.Exit(-1)
	}

	intUserGuess,err := strconv.ParseInt(userGuess)
	if err != nil {
		pl("Incorrect input")
		os.Exit(-2)
	} 

	var numOfTries int = 1;
	for true{
		userGuess,err := reader.ReadString('\n')
		if err != nil {
			os.Exit(-1)
		}
		userGuess = strings.TrimSpace(userGuess)
		if(userGuess == "Enough"){
			pl("Alright, the number was:", randNum)
			break;
		}
		intUserGuess,err := strconv.ParseInt(userGuess)
		if err != nil {
			pl("You misstyped something!")
		} else {
			if intUserGuess == randNum{
				pl("Correct, you have guessed it in:", numOfTries, "tries!")
				break;
			} else if intUserGuess > randNum {
				pl("Lower")
			} else {
				pl("Higher")
			}
		}
		numOfTries++
	}
}