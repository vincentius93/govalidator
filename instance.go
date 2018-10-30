package validator

import "regexp"

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
	mail := regexp.MustCompile(MAIL_FORMAT)
	if !mail.MatchString(a.val.(string)) && a.required == true{
		return myerr(INVALID_EMAIL,a.name)
	}
	return nil
}