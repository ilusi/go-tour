package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Cannot square root negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
	    return 0, ErrNegativeSqrt(x)
	}
	
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))   // 1.4142135623730951 <nil>
	fmt.Println(Sqrt(-2))  // 0 Cannot square root negative number: -2 
}
