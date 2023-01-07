package main

import "fmt"

func main(){
	var ustr string

	fmt.Print("Enter ant String as an Input : ")
	fmt.Scanln(&ustr)

	reversestr := ""

	for i := len(ustr)-1; i >= 0; i--{
		reversestr += string(ustr[i])
	}

	for i := range(ustr){
		if ustr[i] != reversestr[i]{
			fmt.Println(" Given String is not an palindrome string ")
			break
		}
	}
	fmt.Println(" Given string is an palindrome string ")
}