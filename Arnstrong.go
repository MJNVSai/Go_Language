package main

import "fmt"

func main(){
	var arm int
	fmt.Print(" Enter Any Number : ")
	fmt.Scanln(&arm)

	temp := arm
	sum := 0 

	for arm > 0{
		rem := arm%10
		sum += (rem*rem*rem)
		arm = arm/10
	}

	if sum == temp{
		fmt.Println(" Given Number is an Armstrong Number ")
	}else{
		fmt.Println(" Given Number is not an Armstrong Number ")
	}
}