package main

import ("time"
"fmt"
)

func main(){
	var i int;
	i =2;
	fmt.Print("Write", i,"as")
	switch i{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	}
switch time.Now().Weekday(){
case time.Saturday, time.Sunday:
	fmt.Println("Its a Weekend")
default:
	fmt.Println("Its a weekday")
}

t := time.Now()
switch{
case t.Hour() < 12:
	fmt.Println("Its before noon")
default:
	fmt.Println("Its after noon")
}

}