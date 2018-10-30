# GoValidator


GoValidator is a Package validator implements value validations for structs and nested structs fields based on tags.


## Tags
----
- ##### [`field` ] value : required
- ##### [`type`] value : number / text
- ##### [`min`] value : (string number)
- ##### [`max`] value : (string number)
- ##### [`startswith`] value : (string)
- ##### [`endswith`] value : (string)
- ##### [`format`] value : email


Tag `min` / `max` will automatically validate based on the value type.
String : count the length
Int/numeric : validate the value

## Changelog
### 30-10-2018
- Add Handler for email format when its not required
- Add startswith and Endswith validation
----
### Installation

```
go get github.com/vincentius93/govalidator
```
## Usage
---
#### Basic
`````
type User struct {
	FirstName      string       `field:"required"`
	LastName       string       `type:"text"`
	Age            int          `type:"number"`
	Email          string       `format:"email"`
}

func main(){
    MyStruct := User{FirstName:"John",LastName:"S"}
    err := validator.validate(MyStruct)
    if err != nil {
        fmt.Println(err)
    }
}
`````
#### Multiple validation
`````
type User struct {
	FirstName      string       `field:"required" type:"text" max:"10"`
	LastName       string       `field:"required" type:"text" min:"1"`
	Age            int          `type:"number" max:"1"`
	Email          string       `format:"email"`
}

func main(){
    MyStruct := User{FirstName:"John",LastName:"S"}
    err := validator.validate(MyStruct)
    if err != nil {
        fmt.Println(err)
    }
}
`````
#### Nested struct validation
`````
type User struct {
	FirstName      string       `field:"required" type:"text" max:"10"`
	LastName       string       `field:"required" type:"text" min:"1"`
	Age            int          `type:"number" max:"1"`
	Email          string       `format:"email"`
}
type Orders struct {
	OrderId             int  `field:"required" max:"10"`
	UserData            User
}

func main(){
    MyStruct := User{FirstName:"John",LastName:"S"}
    MyOrder := Orders{OrderId:9930438,UserData:MyStruct}

    err := validator.validate(MyOrder)
    if err != nil {
        fmt.Println(err)
    }
}
`````
