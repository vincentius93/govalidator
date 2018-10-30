package main

import (
	"fmt"
	"govalidator"
)

// User contains user information
type User struct {
	FirstName      string
	LastName       string
	Age            uint8
	Email          string
	FavouriteColor string
	Myage			int		   `min:"1" max:"10"`
	Address			Address
}

// Address houses a users address information
type Address struct {
	Street string `format:"email" type:"text"`
	City   string
	Planet string `startswith:"u" endswith:"s"`
	Phone  string `min:"1" max:"10"`
}

func main() {
	fmt.Println("RETURN VALUE OF SIMPLE STRUCT")
	simpleStruct()
	fmt.Println("RETURN VALUE OF SIMPLE NESTEDSTRUCT")
	nestedSturct()
}


func simpleStruct(){
	d := Address{City:"BANDUNG",Phone:"1234",Street:"aas@gmail.com",Planet:"uranusa"}
	err := validator.Validate(d)
	fmt.Println(err)
}

func nestedSturct(){
	d := Address{City:"BANDUNG",Phone:"88766524928347384",Street:"asd@gmail.com"}
	a := User{FirstName:"udin",Age:0,Myage:3,Address:d}
	err := validator.Validate(a)
	fmt.Println(err)
}