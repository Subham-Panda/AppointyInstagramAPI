package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

func CompareJSONToStruct(bytes []byte, empty interface{}) bool {
	var mapped map[string]interface{}

	if err := json.Unmarshal(bytes, &mapped); err != nil {
		return false
	}

	emptyValue := reflect.ValueOf(empty).Type()

	if len(mapped) != emptyValue.NumField() {
		return false
	}

	for key := range mapped {
		if field, found := emptyValue.FieldByName(key); found {
			if !strings.EqualFold(key, strings.Split(field.Tag.Get("json"), ",")[0]) {
				return false
			}
		}
	}

	return true
}