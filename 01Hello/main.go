package main

import (
	"fmt"
	"math"
)

const s string ="constant";
func main(){
	// fmt.Println("Hello");
	// var a = "String";
	// fmt.Println(a);

	// var b,c int = 1,2;
	// fmt.Println(b,c);

	// var d bool = true;
	// fmt.Println(d);

	// var e bool;
	// fmt.Println(e);

	// f := "Apple";
	// fmt.Println(f)

	//   Constant : Go supports constant of character, string, boolean, and numeric values

	// const declares a constant value. and it can appear anywhere a var statement can.

	fmt.Println(s);
	const n = 500000;
	const d = 3e20 /n;
	fmt.Println(d);
	fmt.Println(int64(d));

	fmt.Println(math.Sin(n))

}