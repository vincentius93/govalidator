package validator

var tags =[...]string{
	"field",//required
	"type", //number, text
	"min", //number of min value
	"max", // number of max value
	"length", // number of length value
	"min_length", //number of minimum length value
	"format",//email
}

const MAIL_FORMAT  = `\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`

type myvalidator interface {
	validate()(error)
}
type structDetail struct {
	name 		string
	val			interface{}
	tag_name	string
	tag_value	string
}
type num_error int
type def struct {*structDetail}
type number struct {*structDetail}
type text struct {*structDetail}
type email struct {*structDetail}
type required struct {*structDetail}
type min struct {
	min int
	detail 		*structDetail
}
type max struct {
	max int
	detail 		*structDetail
}
