package main

import "fmt"
import "strings"

func main(){
	x := 10
	y := 20

	z := simple_add(x,y)
	fmt.Println(" Addition : ", z)

	u1, u2, u3 := multiple_return()
	fmt.Println(" Book Number : ", u1)
	fmt.Println(" Book Name : ", u2)
	fmt.Println(" Book Cost : ", u3)

	// Anonymous Function
	Anon := func(a, b, c int) int{
		fmt.Println()
		fmt.Println(" This Is an Anonymous Function")
		return a+b+c
	}
	fmt.Println("Arithmetic : ", Anon(1,2,3))

	cs := vardiac_con("VRSEC", "It", "Cse", "Ece", "Civil", "Mech")
	fmt.Println("Concatination Of Strings : ", cs)

	fact := recursive_fact(10)
	fmt.Println("Factorial of 10 is : ", fact)

	defer defer_end()

	r1 := special(10, inc)
	r2 := special(100, dec)
	fmt.Println("r1 value : ", r1);
	fmt.Println("r2 value : ", r2);

	
}

func simple_add(a int, b int) int{
	fmt.Println(" This Is a Simple Function ")
	ans := a + b

	return ans
}

func multiple_return() (int, string, float64){
	fmt.Println()
	fmt.Println("This function Returns Multiple Values at a Same Time")

	book_no := 38
	book_name := "Simulated Reality"
	book_cost := 480.95

	return book_no, book_name, book_cost
} 

func vardiac_con(elements ...string) string{
	fmt.Println()
	fmt.Println(" This is an vardiac Function ")

	concat := strings.Join(elements, " $ ")
	return concat
}

func recursive_fact(a int) int{
	fmt.Println()
	fmt.Println(" This is an Recursive Function")

	if a == 0 || a == 1{
		return 1
	} else{
		return a*recursive_fact(a-1)
	}
} 

func defer_end() {
	fmt.Println()
	fmt.Println(" This is a Defer Function Call ")

	fmt.Println("After The Main Method The Defer Statement Will Execute")
}

func inc(x int) int{
	x++
	return x
}

func dec(x int) int{
	x--
	return x
}

func special(x int, f func(int) int) int{
	fmt.Println()
	fmt.Println(" This is a Function as a parameter to Another function ")
	r := f(x)
	return r
}

/*

func name(variable datatype, variable datatype) function_return_type{
	Body of the Function
	... => means Ellipsis operator
}

*/