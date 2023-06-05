package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func (v structValidate) max()error{

	switch v.ValueType {
	case reflect.String:
		tagVal ,err := strconv.Atoi(v.TagVAlue)
		if err != nil {
			return err
		}
		if len(v.Value.(string)) > tagVal{
			return errors.New(fmt.Sprintf(invalidMaxValueString.String(),v.TagVAlue,"char"))
		}
	case reflect.Int:
		tagVal ,err := strconv.Atoi(v.TagVAlue)
		if err != nil {
			return err
		}
		if v.Value.(int) > tagVal{
			return errors.New(fmt.Sprintf(invalidMaxValueNumber.String(),v.TagVAlue ))
		}
	case reflect.Float64:
		tagVal ,err := strconv.Atoi(v.TagVAlue)
		if err != nil {
			return err
		}
		if v.Value.(float64) > float64(tagVal){
			return errors.New(fmt.Sprintf(invalidMaxValueNumber.String(),v.TagVAlue ))
		}
	}

	return nil
}
func (v structValidate) min()error{

	switch v.ValueType {
	case reflect.String:
		tagVal ,err := strconv.Atoi(v.TagVAlue)
		if err != nil {
			return err
		}
		if len(v.Value.(string)) < tagVal{
			return errors.New(fmt.Sprintf(invalidMinValueString.String(), v.TagVAlue, "char"))
		}
	case reflect.Int:
		tagVal ,err := strconv.Atoi(v.TagVAlue)
		if err != nil {
			return err
		}
		if v.Value.(int) < tagVal{
			return errors.New(fmt.Sprintf(invalidMinValueNumber.String(), v.TagVAlue))
		}
	case reflect.Float64:
		tagVal ,err := strconv.Atoi(v.TagVAlue)
		if err != nil {
			return err
		}
		if v.Value.(float64) < float64(tagVal){
			return errors.New(fmt.Sprintf(invalidMinValueNumber.String(), v.TagVAlue))
		}
	}

	return nil
}

func (v structValidate) field()error{

	if v.TagVAlue != required{
		return nil
	}
	switch v.ValueType {
	case reflect.String:
		val := strings.TrimSpace(v.Value.(string))
		if val ==""{
			return errors.New(isRequiredString.String())
		}
	case reflect.Int32:
		if v.Value.(int32) == 0 {
			return errors.New(isRequiredNumber.String())
		}
	case reflect.Float64:
		if v.Value.(float64) == 0 {
			return errors.New(isRequiredNumber.String())
		}
	case reflect.Int:
		if v.Value.(int) == 0 {
			return errors.New(isRequiredNumber.String())
		}
	case reflect.Slice:
		if v.BasedValue.Len() == 0{
			return errors.New(isRequiredSlice.String())
		}
	case reflect.Map:
		if len(v.BasedValue.MapKeys()) == 0{
			return errors.New(isRequiredSlice.String())
		}
	}

	return nil
}

func (v structValidate) startsWith() error{
	if v.ValueType != reflect.String{
		return errors.New(invalidTextType.String())
	}
	if !strings.HasPrefix(v.Value.(string), v.TagVAlue) && v.Value.(string)!= ""{
		return errors.New(fmt.Sprintf(invalidStartsWith.String(), v.TagVAlue + "\""))
	}

	return nil
}

func (v structValidate) endsWith() error{

	if v.ValueType != reflect.String{
		return errors.New(invalidTextType.String())
	}

	if !strings.HasSuffix(v.Value.(string), v.TagVAlue) && v.Value.(string)!= ""{
		return errors.New(fmt.Sprintf(invalidEndsWith.String(), v.TagVAlue + "\""))
	}
	return nil
}

func (v structValidate) valueOf() error{

	if v.ValueType != reflect.String{
		return errors.New(invalidTextType.String())
	}
	if v.Value.(string) != ""{
		value := strings.Split(v.TagVAlue,",")
		for i:= 0 ; i < len(value); i++{
			if v.Value.(string) == value[i]{
				return nil
			}
		}
		return errors.New(fmt.Sprintf(invalidValueOf.String(),v.TagVAlue))
	}

	return nil
}

func (v structValidate) format() error{

	if v.ValueType != reflect.String{
		return errors.New(invalidTextType.String())
	}
	if v.Value.(string) != ""{
		switch v.TagVAlue {
		case email:
			mail := regexp.MustCompile(mail_format)
			if !mail.MatchString(v.Value.(string)){
				return errors.New(invalidEmailFormat.String())
			}
		case alphanumeric:
			mail := regexp.MustCompile(alphanumericFormat)
			if !mail.MatchString(v.Value.(string)){
				return errors.New(invalidAlphanumericFormat.String())
			}
		case alphabet:
			mail := regexp.MustCompile(alphabetFormat)
			if !mail.MatchString(v.Value.(string)){
				return errors.New(invalidAlphabetFormat.String())
			}
		}
	}
	return nil
}

func (v structValidate) fieldType() error{

	switch v.TagVAlue {
	case text:
		_,ok := v.Value.(string)
		if !ok {
			return errors.New(invalidTextType.String())
		}
	case number:

		if v.ValueType == reflect.String{
			mail := regexp.MustCompile(numeric)
			if !mail.MatchString(v.Value.(string)){
				return errors.New(invalidNumberType.String())
			}

			return nil
		}

		_,okInt := v.Value.(int)
		_,okFloat := v.Value.(float64)
		if !okInt && !okFloat{

			return errors.New(invalidNumberType.String())
		}
	}
	return nil
}

func (v structValidate) match() error{

	if v.ValueType != reflect.String{
		return errors.New(invalidTextType.String())
	}
	if v.Value.(string) != ""{
		mail := regexp.MustCompile(v.TagVAlue)
		if !mail.MatchString(v.Value.(string)){
			return errors.New(invalidRegexFormat.String())
		}
	}

	return nil
}


func (v validateType) validate(val reflect.Value, tagValue string)error{

	dataType := reflect.ValueOf(val.Interface())
	var validate validateInterface =
		&structValidate{TagVAlue: tagValue, Value: val.Interface(), ValueType: dataType.Kind(),BasedValue: val}

	switch v {
	case min:
		return validate.min()
	case max:
		return validate.max()
	case field:
		return validate.field()
	case startsWith:
		return validate.startsWith()
	case endsWith:
		return validate.endsWith()
	case value_of:
		return validate.valueOf()
	case format:
		return validate.format()
	case fieldType:
		return validate.fieldType()
	case match:
		return validate.match()
	}
	return nil
}