package main

import "fmt"

func main() {

    m := make(map[string]int) //creating a map with key value pair

    // Adding Key value Pairs to Map 
    m["Age"] = 20
    m["Roll"] = 100

    fmt.Println("Initial Map :", m)

    v1 := m["Age"]
    fmt.Println("Value of Age in Map : ", v1)

    fmt.Println("len of an Map :", len(m))

    delete(m, "Roll")
    fmt.Println("Deleting Roll key value pair in map :", m)

    _, prs := m["Roll"]
    fmt.Println("Roll Value :", prs)

    // Another way of creating a map
    n := map[string]int{"EmpID": 45, "Age": 25}
    fmt.Println("Employee map:", n)
}