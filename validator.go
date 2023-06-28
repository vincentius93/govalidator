package validator

import (
	"errors"
	"reflect"
)

func Validate(data interface{})error{
	field := reflect.TypeOf(data)
	value := reflect.ValueOf(data)

	if field.Kind() != reflect.Struct && field.Kind() != reflect.Slice {
		return errors.New(invalidStruct.String())
	}
	if field.Kind() == reflect.Slice{
		err := fetchSliceType(value)
		if err != nil {
			return err
		}
	}

	for a := 0; a < field.NumField(); a ++{
		dataField := field.Field(a)
		kind := dataField.Type.Kind()
		dataValue := value.Field(a)
		if dataValue.IsZero() {
			continue
		}
		switch kind {
		case reflect.Map, reflect.Slice:
			err := fetchMapSlice(dataField,value.Field(a))
			if err != nil {
				return err
			}
			continue
		case reflect.Struct:
			if !dataValue.CanInterface(){
				continue
			}
			err :=  Validate(dataValue.Interface())
			if err != nil {
				return err
			}
			continue
		case reflect.Ptr:
			dataValue = dataValue.Elem()
			switch dataValue.Kind() {
			case reflect.Struct:
				err :=  Validate(value.Field(a).Elem().Interface())
				if err != nil {
					return err
				}
			}
		}
		if !dataValue.CanInterface(){
			continue
		}
		err :=  searchAndValidate(dataField, dataValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func searchAndValidate(data reflect.StructField, value reflect.Value)error{

	for b := allConst-1 ; b >= 0; b -- {
		tagValue,status := data.Tag.Lookup(b.getString())
		if status && tagValue != "" {
			err := b.validate(value,tagValue)
			if err != nil {
				return errors.New(data.Name+" "+err.Error())
			}
		}
	}
	return nil
}

func fetchMapSlice(data reflect.StructField, value reflect.Value)error{

	tagValue,status := data.Tag.Lookup(field.getString())
	if status {
		err := field.validate(value, tagValue)
		if err != nil {
			return err
		}
	}

	switch data.Type.Kind() {
	case reflect.Map:
		for _,v := range value.MapKeys(){
			err := searchAndValidate(data, value.MapIndex(v))
			if err != nil {
				return err
			}
		}
	case reflect.Slice:
		for a := 0 ; a < value.Len(); a++{
			if value.Index(a).Kind() == reflect.Struct{
				err :=  Validate(value.Index(a).Interface())
				if err != nil {
					return err
				}
				continue
			}
			err := searchAndValidate(data, value.Index(a))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func fetchSliceType(data reflect.Value)error{

	for a := 0 ; a < data.Len(); a++{
		err := Validate(data.Index(a).Interface())
		if err != nil {
			return err
		}
	}
	return nil
}