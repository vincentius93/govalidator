package validator

type num_error int

const(
	invalidStruct num_error = iota
	isRequiredString
	isRequiredNumber
	isRequiredSlice
	invalidNumberType
	invalidTextType
	invalidEmailFormat
	invalidAlphabetFormat
	invalidAlphanumericFormat
	invalidMinValueString
	invalidMinValueNumber
	invalidMaxValueString
	invalidMaxValueNumber
	invalidRegexFormat
	invalidStartsWith
	invalidEndsWith
	invalidValueOf

)
func (num num_error)String()string{
	error := [...]string{
		"Not a struct type or memory address of struct!",
		"value is required!",
		"value is required, minimum 1!",
		"value is required, minimum 1 value!",
		"invalid type of number",
		"invalid type of text",
		"value format is not email",
		"value format is not alphabet",
		"value format is not alphanumeric",
		"should be greather than %s %s",
		"should be greather than %s",
		"should be lower than %s %s",
		"should be lower chan %s",
		"invalid value of expression",
		"value should be starts with %s \"",
		"value should be ends with $s \"",
		"value should be contains one of %s",
	}
	return error[num]
}
