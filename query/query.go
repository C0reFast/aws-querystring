package query

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func BindQuery(values url.Values, output interface{}) error {
	outputVal := reflect.ValueOf(output)
	if outputVal.Kind() != reflect.Ptr {
		return fmt.Errorf("output must be a pointer")
	}

	outputElem := outputVal.Elem()
	outputType := outputElem.Type()

	for i := 0; i < outputType.NumField(); i++ {
		field := outputType.Field(i)
		tag := field.Tag.Get("query")
		if tag == "" {
			continue
		}

		value := values.Get(tag)
		fieldVal := outputElem.FieldByName(field.Name)

		if field.Type.Kind() == reflect.Slice {
			elemType := field.Type.Elem()
			if elemType.Kind() != reflect.Struct {
				prefix := tag + "."
				arrIndex := 1
				for {
					currKey := prefix + fmt.Sprint(arrIndex)
					currValue := values.Get(currKey)
					if currValue == "" {
						break
					}

					currSliceVal := reflect.ValueOf(currValue)
					fieldVal.Set(reflect.Append(fieldVal, currSliceVal))
					arrIndex++
				}
			} else {
				prefix := tag + "."
				objIndex := 1
				outer := true
				for outer {
					innerValues := make(url.Values)
					for innerKey, innerValue := range values {
						if strings.HasPrefix(innerKey, prefix+strconv.Itoa(objIndex)+".") {
							innerValues.Set(strings.TrimPrefix(innerKey, prefix+strconv.Itoa(objIndex)+"."), innerValue[0])
						}
					}
					if len(innerValues) == 0 {
						break
					}

					newStructPtr := reflect.New(elemType)
					err := BindQuery(innerValues, newStructPtr.Interface())
					if err != nil {
						return err
					}
					fieldVal.Set(reflect.Append(fieldVal, newStructPtr.Elem()))

					objIndex++
				}
			}
		} else {
			fieldVal.Set(reflect.ValueOf(value))
		}
	}

	return nil
}
