package validator

import (
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
	val,ok :=t.detail.val.(string)
	if ok == true{
		if len(val) < t.min{
			return myerr(MIN,t.detail.name,t.detail.tag_value)
		}
	}else{

		intVal,_ := t.detail.val.(int)
		if intVal < t.min{
			return myerr(MIN,t.detail.name,t.detail.tag_value)
		}
	}
	return nil
}
func (t max)validate()error{
	val,ok :=t.detail.val.(string)
	if ok == true{
		if len(val) > t.max{
			return myerr(MAX,t.detail.name,t.detail.tag_value)
		}
	}else{

		intVal,_ := t.detail.val.(int)
		if intVal > t.max{
			return myerr(MAX,t.detail.name,t.detail.tag_value)
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
