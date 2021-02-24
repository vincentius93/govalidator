package validator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func (d def)validate()error{
	return myerr(INVALID_TYPE,d.name)
}
func (r required)validate()error{
	value_kind := reflect.TypeOf(r.val).Kind()
	value_len := 1
	if value_kind != reflect.Int && value_kind != reflect.Float64{
		value_len = reflect.ValueOf(r.val).Len()
		//r.val = strings.TrimSpace(r.val.(string))
	}
	if value_kind == reflect.String {
		r.val = strings.TrimSpace(r.val.(string))
	}
	if r.val =="" || r.val == nil || value_len == 0 {
		return myerr(ISREQUIRED,parseName(*r.structDetail))
	}
	return nil
}
func (t text)validate()error{
	_,ok :=t.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,parseName(*t.structDetail))
	}
	return nil
}
func (t number)validate()error{
	_,ok :=t.val.(int)
	if ok == false{
		_,validString := t.val.(string)
		if validString == false{
			return myerr(INVALID_INT, t.name)
		}
		_,err := strconv.Atoi(t.val.(string))
		if err != nil {
			return myerr(INVALID_CONVERSION_INT, parseName(*t.structDetail))
		}
	}
	return nil
}
func (t min)validate()error{
	var structVal *structDetail
	structVal = t.detail
	val,ok :=t.detail.val.(string)
	if ok == true{
		if len(val) < t.min{
			return myerr(MIN,parseName(*t.detail),t.detail.tag_value)
		}
	}else{
		dataType :=reflect.ValueOf(t.detail.val).Kind()
		intVal,_ := t.detail.val.(int)
		switch dataType {
		case reflect.Float64:
			return floatValue(*structVal,t.min,"min")
		case reflect.Int32:
			return int32value(*structVal,t.min,"min")
		default:
			if intVal < t.min{
				return myerr(MIN,parseName(*t.detail),t.detail.tag_value)
			}
		}
	}
	return nil
}
func (t max)validate()error{
	var structVal *structDetail
	structVal = t.detail
	val,ok :=t.detail.val.(string)
	if ok == true{
		if len(val) > t.max{
			return myerr(MAX,parseName(*t.detail),t.detail.tag_value)
		}
	}else{
		dataType :=reflect.ValueOf(t.detail.val).Kind()
		intVal,_ := t.detail.val.(int)
		switch dataType {
		case reflect.Float64:
			return floatValue(*structVal,t.max,"max")
		case reflect.Int32:
			return int32value(*structVal,t.max,"max")
		default:
			if intVal > t.max{
				return myerr(MAX,parseName(*t.detail),t.detail.tag_value)
			}
		}
	}
	return nil
}
func (a email)validate()(error){
	_,ok :=a.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,parseName(*a.structDetail))
	}
	if a.required == false && a.val.(string) == ""{
		return nil
	}
	mail := regexp.MustCompile(MAIL_FORMAT)
	if !mail.MatchString(a.val.(string)){
		return myerr(INVALID_EMAIL,parseName(*a.structDetail))
	}
	return nil
}
func (a alphabet)validate()(error){
	_,ok :=a.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,parseName(*a.structDetail))
	}
	if a.required == false && a.val.(string) == ""{
		return nil
	}
	mail := regexp.MustCompile(ALPHABET)
	if !mail.MatchString(a.val.(string)){
		return myerr(INVALID_ALPHABET,parseName(*a.structDetail))
	}
	return nil
}
func (a alphanumeric)validate()(error){
	_,ok :=a.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,parseName(*a.structDetail))
	}
	if a.required == false && a.val.(string) == ""{
		return nil
	}
	mail := regexp.MustCompile(ALPHANUMERIC)
	if !mail.MatchString(a.val.(string)){
		return myerr(INVALID_ALPHANUMERIC,parseName(*a.structDetail))
	}
	return nil
}
func (s startswith)validate()(error){
	_,ok :=s.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,parseName(*s.structDetail))
	}
	if s.required == false && s.val.(string) == ""{
		return nil
	}
	if strings.HasPrefix(s.val.(string), s.tag_value)==false {
		return myerr(STARTS_WITH,parseName(*s.structDetail),s.tag_value)
	}

	return nil
}
func (s endswith)validate()(error){
	_,ok :=s.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,parseName(*s.structDetail))
	}
	if s.required == false && s.val.(string) == ""{
		return nil
	}
	if strings.HasSuffix(s.val.(string), s.tag_value)==false {
		return myerr(ENDS_WITH,parseName(*s.structDetail),s.tag_value)
	}

	return nil
}
func (v valueOf)validate()(error){
	value := strings.Split(v.tag_value,",")
	for i:= 0 ; i < len(value); i++{
		if v.val.(string) == value[i]{
			return nil
		}
	}
	return myerr(VALUE_OF,v.name,v.tag_value)
}
func floatValue(value structDetail,compare int,tipe string)(err error){
	floatVal := value.val.(float64)
	compareVal := float64(compare)
	switch tipe {
	case "min":
		if floatVal < compareVal{
			return myerr(MIN,parseName(value),value.tag_value)
		}
	case "max":
		if floatVal > compareVal{
			return myerr(MAX,parseName(value),value.tag_value)
		}
	}
	return nil
}
func int32value(value structDetail,compare int,tipe string)(err error){
	int32val := value.val.(int32)
	compareVal := int32(compare)
	switch tipe {
	case "min":
		if int32val < compareVal{
			return myerr(MIN,parseName(value),value.tag_value)
		}
	case "max":
		if int32val > compareVal{
			return myerr(MAX,parseName(value),value.tag_value)
		}
	}
	return nil
}
func (v match)validate()(error){
	pattern,_ := regexp.Compile(v.tag_value)
	_,ok:= v.val.(string)
	if !ok{
		return myerr(INVALID_STRING,parseName(*v.structDetail))
	}
	ismatch := pattern.MatchString(v.val.(string))
	if !ismatch{
		return myerr(MATCHSTRING,parseName(*v.structDetail))
	}
	return nil
}