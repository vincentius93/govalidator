package validator

import (
	"reflect"
	"regexp"
	"strings"
)

func (d def)validate()error{
	return myerr(INVALID_TYPE,d.name)
}
func (r required)validate()error{
	if r.val =="" || r.val == nil{
		return myerr(ISREQUIRED,r.name)
	}
	return nil
}
func (t text)validate()error{
	_,ok :=t.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,t.name)
	}
	return nil
}
func (t number)validate()error{
	_,ok :=t.val.(int)
	if ok == false{
		return myerr(INVALID_INT, t.name)
	}
	return nil
}
func (t min)validate()error{
	var structVal *structDetail
	structVal = t.detail
	val,ok :=t.detail.val.(string)
	if ok == true{
		if len(val) < t.min{
			return myerr(MIN,t.detail.name,t.detail.tag_value)
		}
	}else{
		dataType :=reflect.ValueOf(t.detail.val).Kind()
		intVal,_ := t.detail.val.(int)
		switch dataType {
		case reflect.Float32:
			return floatValue(*structVal,t.min,"min")
		default:
			if intVal < t.min{
				return myerr(MIN,t.detail.name,t.detail.tag_value)
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
			return myerr(MAX,t.detail.name,t.detail.tag_value)
		}
	}else{
		dataType :=reflect.ValueOf(t.detail.val).Kind()
		intVal,_ := t.detail.val.(int)
		switch dataType {
		case reflect.Float32:
			return floatValue(*structVal,t.max,"max")
		default:
			if intVal > t.max{
				return myerr(MAX,t.detail.name,t.detail.tag_value)
			}
		}
	}
	return nil
}
func (a email)validate()(error){
	_,ok :=a.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,a.name)
	}
	if a.required == false && a.val.(string) == ""{
		return nil
	}
	mail := regexp.MustCompile(MAIL_FORMAT)
	if !mail.MatchString(a.val.(string)){
		return myerr(INVALID_EMAIL,a.name)
	}
	return nil
}
func (s startswith)validate()(error){
	_,ok :=s.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,s.name)
	}
	if s.required == false && s.val.(string) == ""{
		return nil
	}
	if strings.HasPrefix(s.val.(string), s.tag_value)==false {
		return myerr(STARTS_WITH,s.name,s.tag_value)
	}

	return nil
}
func (s endswith)validate()(error){
	_,ok :=s.val.(string)
	if ok == false{
		return myerr(INVALID_STRING,s.name)
	}
	if s.required == false && s.val.(string) == ""{
		return nil
	}
	if strings.HasSuffix(s.val.(string), s.tag_value)==false {
		return myerr(ENDS_WITH,s.name,s.tag_value)
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

	floatVal := value.val.(float32)
	compareVal := float32(compare)
	switch tipe {
	case "min":
		if floatVal < compareVal{
			return myerr(MIN,value.name,value.tag_value)
		}
	case "max":
		if floatVal > compareVal{
			return myerr(MAX,value.name,value.tag_value)
		}
	}
	return nil
}