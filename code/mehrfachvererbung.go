type A struct {
    a int
    x int
}

type B struct {
    b int
    x int
}

type C struct {
    A
    B
}

func main() {
    var c C
    c.a = 1     // entspricht c.A.a
    c.b = 2     // entspricht c.B.b

    c.A.x = 4   // keine Kurzschreibweise m�glich
    c.B.x = 5   // keine Kurzschreibweise m�glich
}