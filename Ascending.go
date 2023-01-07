package main

import "fmt"
import "sort"

func main(){
	array := [6]int{}

	fmt.Println(" Enter the elements into the array ")
	for i := 0; i < 6; i++{
	  fmt.Printf("Enter %vth element: ", i)
      fmt.Scanln(&array[i])
	}

	fmt.Println("\nInput array is : ", array)

	sort.Ints(array[:])
	fmt.Println(" Sorted array is : ", array)
}