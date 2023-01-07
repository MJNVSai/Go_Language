package main

import ("fmt")

func main(){

	var arr1 = [3]int{1,2,3} // 3 indicates the length of an array
	arr2 := [5]int{4,5,6,7,8}

	arr3 := [...]string{"Volvo", "BMW", "Mercendes", "Aadhipurush", "SkyScrapper"} // [...] => indicates unlimited length in the array

	fmt.Printf("Type of an array : %T, \n Values in the array are : %v", arr1, arr1)
	fmt.Printf("Type of an array : %T, \n Values in the array are : %v", arr3, arr3)

	// Performing the Slicing Operations
	fmt.Println("\nElement present at the position 3 is in arr2 variable : ", arr2[3])

	// Manipulating the elements in the array
	fmt.Println("List of all Elements in arr2 are : ", arr2)
	fmt.Println("Element present in position 2 is : ", arr2[2])
	arr2[2] = 500
	fmt.Println("Now Element present in position 2 is : ", arr2[2])

	// Special Type in Array Declaration
	array := [...]int{50:23, 12:100} // i.e; element 23 is placed at 50th position in array
	// element 100 is placed at 12th position in the array
	fmt.Printf("Type of an array : %T, \n Values in the array are : %v", array, array)

	fmt.Println("\nLength of an array is : ", len(array))
	fmt.Println("\n Capacity of an array is : ", cap(array))

	// 2 dimensional Array
	TwoD := [3][3]int{ {1,2,3}, {4,5,6}, {7,8,9} }
	fmt.Printf("\n Type of an 2D array is : %T ", TwoD)
	fmt.Println("\n Elements in 2D array are : ", TwoD)

}