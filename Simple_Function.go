package main

import "fmt"

func Simple(x int, y int) int{
	fmt.Println(" This is an Simple Function")
	return x+y
}

func main(){
	result := Simple(10,99)
	fmt.Println(" Addition of 2 Numbers : ", result)
}