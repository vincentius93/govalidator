package validator

import "reflect"

type validateType int

const (
	max validateType = iota
	min
	field
	startsWith
	endsWith
	value_of
	format
	fieldType
	match
	allConst
)

func (v validateType)getString()string{
	data := []string{
		"max",
		"min",
		"field",
		"startswith",
		"endswith",
		"value_of",
		"format",
		"fieldType",
		"match",
	}
	return data[v]
}

const (
	mail_format  = `\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`
	alphanumericFormat = `^[A-Za-z0-9 ]*$`
	alphabetFormat = `^[A-Za-z ]*$`
	numeric = `^[0-9]*$`
	email = "email"
	alphanumeric = "alphanumeric"
	required = "required"
	number = "number"
	text = "text"
	alphabet = "alphabet"
)

type structValidate struct {
	TagVAlue  string
	Value     interface{}
	BasedValue reflect.Value
	ValueType reflect.Kind
}

type validateInterface interface {
	fieldType() error
	format() error
	min() error
	max() error
	valueOf() error
	startsWith() error
	endsWith() error
	field() error
	match() error
}
