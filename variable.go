package main

import "fmt"

func main(){
	fmt.Println("Different types of variable declarations.")

	var k int
	k = 10
	// m = 10 is a wrong format
	m := 10
	fmt.Println("k and m values are : ", k, " ", m)

	var b, c int = 1, 3
	var k1 = 34
	var K1 = 32
	fmt.Println("values of b, c, k1, K1 are : ", b, " ", c, " ", k1, " ", K1)
}