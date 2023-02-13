package main

import(
	"fmt"
	"math/rand"
	"time"
	"math"
)

var pl = fmt.Println

func main(){
	pl("5+4 = ", 5+4)
	pl("5-4 = ", 5-4)
	pl("5*4 = ", 5*4)
	pl("5/4 = ", 5/4)
	pl("5%4 = ", 5%4)

	mInt := 1
	mInt += 1
	mInt++
	pl("Should be 3 =", mInt)

	//in order to generate truely random numbers with each run
	//we need the seed generation below
	seedSecs := time.Now().UnixNano()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50)+1
	pl("Random :", randNum)
	pl("Abs(-10) ", math.Abs(-10))
	pl("Pow(4, 2) ", math.Pow(4, 2))
	pl("Sqrt(16) ", math.Sqrt(16))
	pl("Round(4.4) ", math.Round(4.4))
	pl("Log2(8) ", math.Log2(8))
	pl("Max(5,4) ", math.Max(5,4))
	pl("Min(5,4) ", math.Min(5,4))
	//Cos, Tan, Acos, Asin
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 * math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	pl("Sin(90) ", math.Sin(r90))
}