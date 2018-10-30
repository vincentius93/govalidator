package validator

import (
	"reflect"
	"strconv"
)

func Validate(mystruct interface{})(err error){

	field := reflect.TypeOf(mystruct)
	value := reflect.ValueOf(mystruct)
	if field.Kind() != reflect.Struct{
		return myerr(STRUCT_ERROR)
	}
	for i:= 0 ; i< field.NumField();i++{
		kind := field.Field(i).Type.Kind()
		val := value.Field(i).Interface()
		if kind == reflect.Struct{
			err = Validate(val)
		}else if kind != reflect.Map || kind != reflect.Array{
			err = validateTags(field.Field(i),val)
		}
		if err != nil {
			return
		}
	}
	return err
}
func validateTags(field reflect.StructField,value interface{})(err error){

	for i := 0; i< len(tags);i++{
		tag,stat := field.Tag.Lookup(tags[i])
		field_tag,stat := field.Tag.Lookup("field")
		required := false
		if stat == true{
			if stat == true && field_tag == "required"{
				required = true
			}
			structInfo := structDetail{
				name:field.Name,
				val:value,
				tag_name:tags[i],
				tag_value:tag,
				required:required,
			}
			err:= getType(&structInfo).validate()
			if err != nil {
				return err
			}

		}
	}
	return err
}

func getType(s *structDetail)myvalidator{

	tagsValue := map[string]myvalidator{
		"email":email{s},
		"text":text{s},
		"number":number{s},
		"required":required{s},
		"default":def{s},
	}
	res := tagsValue["default"]
	switch s.tag_name {
	case "min":
		lenChar,err := strconv.Atoi(s.tag_value)
		if err != nil{
			res = tagsValue["default"]
		}
		tagsValue[s.tag_name] = min{min:lenChar,detail:s}
		res = tagsValue["min"]
	case "max":
		lenChar,err := strconv.Atoi(s.tag_value)
		if err != nil{
			res = tagsValue["default"]
		}
		tagsValue[s.tag_name] = max{max:lenChar,detail:s}
		res = tagsValue["max"]
	default:

		if _, ok := tagsValue[s.tag_value]; ok {
			res = tagsValue[s.tag_value]
		}
	}

	return res
}







