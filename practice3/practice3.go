package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 5), pow(3, 3, 20))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
} 

func pow(x, n, lim float64) float64 {
	
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("In the if block")
		return v
	} else {
		fmt.Print("In else block")
		fmt.Printf("The Value of v is %v\n", v)
		return lim
	}
	
}
