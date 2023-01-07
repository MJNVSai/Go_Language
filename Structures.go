package main

import "fmt"

type person struct {
    Name string
    Age  int
}

func newPerson(name string) *person {

    p := person{Name: name}
    p.Age = 42
    return &p
}

func main() {

    fmt.Println("directly give the values to the struct : ", person{"Bob", 20})

    fmt.Println("giving the values by mapping the variable names in the struct : ", person{Name: "Alice", Age: 30})

    fmt.Println("giving only one value to the struct : ", person{Name: "Fred"})

    fmt.Println("Another way of giving values to the struct : ", &person{Name: "Ann", Age: 40})

    fmt.Println("New Person Function : ", newPerson("Jon"))

    s := person{Name: "Sean", Age: 50}
    fmt.Println(" Initial Structure : ", s)
    fmt.Println("Getting the Name value in struct using object : ", s.Name)

    sp := &s
    fmt.Println("Age of an Employee : " , sp.Age)

    sp.Age = 51
    fmt.Println(" Change in Age value in Struct : ", sp.Age)
}