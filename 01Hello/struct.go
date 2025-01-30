package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}


// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

// Run codeCopy code
// package main
// import "fmt"
// This person struct type has name and age fields.

// type person struct {
//     name string
//     age  int
// }
// newPerson constructs a new person struct with the given name.

// func newPerson(name string) *person {
// Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.

//     p := person{name: name}
//     p.age = 42
//     return &p
// }
// func main() {
// This syntax creates a new struct.

//     fmt.Println(person{"Bob", 20})
// You can name the fields when initializing a struct.

//     fmt.Println(person{name: "Alice", age: 30})
// Omitted fields will be zero-valued.

//     fmt.Println(person{name: "Fred"})
// An & prefix yields a pointer to the struct.

//     fmt.Println(&person{name: "Ann", age: 40})
// It’s idiomatic to encapsulate new struct creation in constructor functions

//     fmt.Println(newPerson("Jon"))
// Access struct fields with a dot.

//     s := person{name: "Sean", age: 50}
//     fmt.Println(s.name)
// You can also use dots with struct pointers - the pointers are automatically dereferenced.

//     sp := &s
//     fmt.Println(sp.age)
// Structs are mutable.

//     sp.age = 51
//     fmt.Println(sp.age)
// If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.

//     dog := struct {
//         name   string
//         isGood bool
//     }{
//         "Rex",
//         true,
//     }
//     fmt.Println(dog)
// }

// package main

// import (
//     "fmt"
//     "testing"
// )

// func IntMin(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// func TestIntMinBasic(t *testing.T) {
//     ans := IntMin(2, -2)
//     if ans != -2 {

//         t.Errorf("IntMin(2, -2) = %d; want -2", ans)
//     }
// }

// func TestIntMinTableDriven(t *testing.T) {
//     var tests = []struct {
//         a, b int
//         want int
//     }{
//         {0, 1, 0},
//         {1, 0, 0},
//         {2, -2, -2},
//         {0, -1, -1},
//         {-1, 0, -1},
//     }

//     for _, tt := range tests {

//         testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
//         t.Run(testname, func(t *testing.T) {
//             ans := IntMin(tt.a, tt.b)
//             if ans != tt.want {
//                 t.Errorf("got %d, want %d", ans, tt.want)
//             }
//         })
//     }
// }

// func BenchmarkIntMin(b *testing.B) {

//     for i := 0; i < b.N; i++ {
//         IntMin(1, 2)
//     }
// }
