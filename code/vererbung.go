package main

import "fmt"

type A struct {
    a1 int
    a2 int
}

type B struct {
    A         // namenlose Einbindung des Verbunds A
    b1 int
    b2 int
}

func main() {
    var b B
    b.a1 = 1   // entspricht b.A.a1
    b.a2 = 2   // entspricht b.A.a2
    b.b1 = 3
    b.b2 = 4
    fmt.Println(b)
}
