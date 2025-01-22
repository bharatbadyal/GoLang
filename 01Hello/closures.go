package main
import "fmt"

// func intSeq() func()int{
// 	i:= 0
// 	return func()int{
// 		i++
// 		return i
// 	}
// func main(){
// 	nextInt := intSeq()

// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())

// 	newInts := intSeq()
//     fmt.Println(newInts())

// }

// func fact(n int) int{
// 	if n==0{
// 		return 1
// 	}
// 	return n*fact(n-1)
// }

// func main(){
// 	fmt.Println(fact(7))

// 	var fib func(n int)int

// 	fib = func(n int)int{
// 		if n<2{
// 			return n
// 		}
// 		return fib(n-1) + fib(n-2)
// 	}

	// fmt.Println(fib(7))

	package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
}