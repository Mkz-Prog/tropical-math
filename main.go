package main

import (
	"github.com/Mkz-Prog/tropical-math/math"
	"fmt"
)

func main() {
	var a,b float64
	fmt.Print("Enter two numbers to add: ")
	fmt.Scanf("%v %v", &a, &b)
	r1, r2 :=  math.Rational{Value: a, IsInfinity: false}, math.Rational{Value: b, IsInfinity: false}

	fmt.Printf("%v is your result!\n", r1.Add(r2).Value)

}



