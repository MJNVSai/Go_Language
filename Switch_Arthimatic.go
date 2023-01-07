package main

import "fmt"
import "os"

func main(){
	var a,b, choice int64
	fmt.Print("Enter the 1st Number : ")
	fmt.Scanln(&a)
	fmt.Print("Enter the 2nd Number : ")
	fmt.Scanln(&b)

	for true{
		fmt.Println(" \nEnter 1 for Addition ")
		fmt.Println(" Enter 2 for Subraction ")
		fmt.Println(" Enter 3 for Multiplication")
		fmt.Println(" Enter 4 for Division")
		fmt.Println(" Enter 5 for to exit the program")

		fmt.Print("\n Enter your Choice : ")
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			fmt.Println(" Addition of Given Numbers are : ", a+b)

		case 2:
			fmt.Println(" Subraction of Given Numbers are : ", a-b)

		case 3:
			fmt.Println(" Multiplication of Given Numbers are : ", a*b)

		case 4:
			fmt.Println(" Division of Given Numbers are : ", a/b)

		case 5:
			os.Exit(0)

		default:
			fmt.Println(" Invalid Input ")
		}
	}
}