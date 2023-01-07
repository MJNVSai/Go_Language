package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/mail"
    "phonenumbers"
    //"github.com/nyaruka/phonenumbers/cmd/buildmetadata"
)

type Login struct {
    Firstname string `Firstname`
    Lastname string `Lastname`
    Mobile int64 `Mobile Number`
    Email string `Email`
}


func Email_Validate(email string){
    // validating the email
    addr, err := mail.ParseAddress(email)

    if err != nil{
        fmt.Println("\n The Email you have entered is Invalid please check it again !! \n")
    }else{
        fmt.Println("\n The Email you have entered is valid and the entered Email is ", ": ", addr.Address, "\n")
    }
}

func main(){
	//fmt.Println(Login{"sai", "venkat", 78456798234, "whoareyou@gmail.com"})

    var firstname, lastname, email string
    var mobile int64

    fmt.Print("Enter FirstName : ")
    fmt.Scanln(&firstname)
    fmt.Print("Enter Lastname : ")
    fmt.Scanln(&lastname)
    fmt.Print("Enter Mobile Number : ") 
    fmt.Scanln(&mobile)
    fmt.Print("Enter Email : ")
    fmt.Scanln(&email)

    sign := Login{firstname, lastname, mobile, email}
    json_login, err := json.Marshal(sign)

    if err != nil {
        log.Fatal(err)
    }

    Email_Validate(email)

    fmt.Println(string(json_login))
    //fmt.Println(json_login)
}

