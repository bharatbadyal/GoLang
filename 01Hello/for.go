package main

import "fmt"

func main(){
	fmt.Println("Hello From For");
	
	for i:=0; i<10 ; i++{
		fmt.Println("i := ",i)
	}

	for n:= range 10{
		if(n%2 ==0){
			fmt.Println("Even no. := ",n);
		}
	}
}

