package main

import "fmt"
import "unicode/utf8"

func main(){
	var intNum int = 32767
	intNum =+ 1;
	fmt.Println(intNum)

	var floNum float32 = 454.3
	fmt.Println(float32(intNum)+ floNum);

	var stringVal string = "Hello" + " " + "World"
	fmt.Println("The given massage is: ", stringVal)
	fmt.Println(len(stringVal))

	var myRune rune = 'a'
	fmt.Println(myRune)

	fmt.Println(utf8.RuneCountInString("yw"))
}