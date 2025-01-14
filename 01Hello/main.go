// package main

// import (
// 	"fmt"
// 	// "math"
// )

// // const s string ="constant";
// func main(){
// 	// fmt.Println("Hello");
// 	// var a = "String";
// 	// fmt.Println(a);

// 	// var b,c int = 1,2;
// 	// fmt.Println(b,c);

// 	// var d bool = true;
// 	// fmt.Println(d);

// 	// var e bool;
// 	// fmt.Println(e);

// 	// f := "Apple";
// 	// fmt.Println(f)

// 	//   Constant : Go supports constant of character, string, boolean, and numeric values

// 	// const declares a constant value. and it can appear anywhere a var statement can.

// 	// fmt.Println(s);
// 	// const n = 500000;
// 	// const d = 3e20 /n;
// 	// fmt.Println(d);
// 	// fmt.Println(int64(d));

// 	// fmt.Println(math.Sin(n))

// 	// Go by examples := for

// 	i := 1;
// 	for i<=3{
// 		fmt.Println(i);
// 		i++;
// 	}

// 	for j:=0; j<3; j++{
// 		fmt.Println(j)
// 	}

// 	for i:= range 3{
// 		fmt.Println("Range", i)
// 	}

// 	for{
// 		fmt.Println("loop")
// 		break
// 	}

// 	for n:= range 6{
// 		if n%2 != 0{
// 			continue
// 		}
// 		fmt.Println(n)
// 	}
// }

package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Welcome to Badlapur</h1>")
}

func main(){
	http.HandleFunc("/",handlerFunc)
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000",nil)
}
