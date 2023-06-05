package main

import (
	"fmt"
	"govalidator"
)

// User contains user information
type User struct {
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name"`
	Age            uint8
	Email          string
	FavouriteColor string
	Address        *Address
	Myage          *int    `min:"1" max:"10"`
	TestPtr        *string `fieldType:"number"`
}

// Address houses a users address information
type Address struct {
	Street        string `format:"email" Fieldtype:"text"`
	City          string `format:"alphanumeric"`
	Planet        string `startswith:"u" endswith:"s"`
	Phone         string `min:"1" max:"10" fieldType:"number"`
	DetailAddress string `match:"[0-9]"`
}
type User1 struct {
	FirstName      string            `value_of:"jhonny,john" field:"required"`
	Age            float64           `json:"umur" min:"4" max:"100" field:"required"`
	FavouriteColor []string          `field:"required"`
	Myage          int               `min:"1" max:"10"`
	Testing        float64           `field:"required"`
	DataMap        map[string]string `fieldType:"number"`
	Address        []Address
	Email          string
	LastName       int
}

func main() {
	//fmt.Println("RETURN VALUE OF SIMPLE STRUCT")
	//simpleStruct()
	//fmt.Println("RETURN VALUE OF SIMPLE NESTEDSTRUCT")
	//nestedSturct()
	//fmt.Println("RETURN VALUE OF ARRAY STRUCT")
	//arrayStruct()
	//fmt.Println("RETURN VALUE OF NESTED ARRAY STRUCT")
	nestedArrayStruct()
	//CustomRegexValidate()
}

func simpleStruct() {
	d := Address{City: "BANDUNG TEST 123", Phone: "9663", Street: "aas@gmail.com", Planet: "uranus"}
	err := validator.Validate(d)
	fmt.Println(err)
}

func nestedSturct() {
	b := "123"
	d := Address{City: "BANDUNG!!", Phone: "8876652494", Street: "asd@gmail.com"}
	a := User{FirstName: "udin", Age: 0, Address: &d,TestPtr: &b}


	err := validator.Validate(a)
	fmt.Println(err)
}

func arrayStruct() {
	var d []Address
	d = append(d, Address{City: "ASD", Phone: "asd", Planet: "uranus", Street: "asd@gmail.com"})
	d = append(d, Address{City: "ASD", Phone: "asd", Planet: "mars", Street: "asd@agmail.com"})
	err := validator.Validate(d)
	fmt.Println(err)
}
func nestedArrayStruct() {
	var d []Address
	d = append(d, Address{City: "ASD", Phone: "asd", Planet: "uranus", Street: "asd@gmail.com"})
	d = append(d, Address{City: "ASD", Phone: "asd", Planet: "uranus", Street: "asd@agmail.com"})
	a := User1{FirstName: "john", Age: 4.5, Myage: 8, FavouriteColor: []string{"asd"}, Address: d}
	a.DataMap = make(map[string]string)
	a.DataMap["keys1"] = "asd123"
	err := validator.Validate(a)
	fmt.Println(err)
}

func CustomRegexValidate() {
	d := Address{City: "BANDUNG", Phone: "1231231231", Street: "asd@gmail.com", DetailAddress: "1", Planet: "uranus"}
	err := validator.Validate(d)
	fmt.Println(err)
}
