package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area of a Rectangle : ", r.area())
    fmt.Println("perimeter of a Rectangle :", r.perim())

    rp := &r
    rp.width = 20 
    rp.height = 50
    fmt.Println("\narea of a Rectangle 2 : ", rp.area())
    fmt.Println("perim of a Rectangle 2 :", rp.perim())
}