//go build hello-world.go ==> it converts .go code to an encrypted binary code language
//dir
//./hello-world ==> this command used for running the file.go

package main

// By using paranthesis () after import we can import multiple inbuilt functions and packages
import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {

    // Concatination Of Strings
    str1 := "go"
    str2 := " language"
    fmt.Println("String 1 : ", str1)
    fmt.Println("String 2 : ", str2)
    fmt.Println("Concatination of 2 Strings : ", str1 + str2)

    fmt.Println()

    // Performing Arthimatic Operations
    num1 := 5
    num2 := 10
    num3 := 5.65
    num4 := 8.9
    fmt.Println("Number 1 : ", num1)
    fmt.Println("Number 2 : ", num2)
    fmt.Println("Addition of 2 Numbers are : ", num1 + num2)
    fmt.Println("Subraction of 2 Numbers are : ", num1 - num2)
    fmt.Println("Multiplication of 2 Numbers are : ", num1 * num2)
    fmt.Println("Division of 2 Float Numbers : ", num3 / num4)

    fmt.Println()
    
    // Performing Logical Operations
    var b1 bool = true
    var b2 bool = false
    fmt.Println("Data Stored in Variable b1 is : ", b1)
    fmt.Println("Data Stored in Variable b2 is : ", b2)
    fmt.Println("Applying AND Opearation with different values : ", b1 && b2)
    fmt.Println("Applying AND Opearation with Same Values : ", b1 && b1)
    fmt.Println("Applying OR Opearation with different values: ", b1 || b2)
    fmt.Println("Applying OR Opearation with Same values: ", b2 || b2)
    fmt.Println("Applying NOT(!) Opearation : ", !b1)
    
    fmt.Println()

    // Performing Constatant Data Type
    fmt.Println("Value Stored in S variable which is a Constatnt DataType : ", s)
    //const n = 50000000 // explicit conversion.
    const n = 50

    const d = 3e20 / n
    fmt.Println("the value of 3e20 is : ", 3e20)

    fmt.Println("Converting the Biggest number : ", int64(d))

    fmt.Println(" value of sin 50 is : ", math.Sin(n)) 
    
    fmt.Println()

    //Performing the Conditional and Control Structures
    i := 1
    fmt.Println("for loop with Single Condition printing 1 to 3 Numbers :")
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    fmt.Println("for loop with 3 conditions printing 7 to 9 Numbers : ")
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    fmt.Println("for loop with no conditions : ")
    for {
        fmt.Println("loop")
        break
    }

    fmt.Println("Printing the Odd numbers using For loop : ")
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
    
    // Checking the even odd Conditions
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // Checking the divisibility Conditions
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // checking the Positive, Negative, Single digits 
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }



}



