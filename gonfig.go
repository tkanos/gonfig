// Package gonfig implements simple configuration reading
// from both JSON files and enviornment variables.

package gonfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"github.com/ghodss/yaml"
)

// tag name to override the field name of an environment variable
const envTagName = "env"

// GetConf aggregates all the JSON and enviornment variable values
// and puts them into the passed interface.
func GetConf(filename string, configuration interface{}) (err error) {

	configValue := reflect.ValueOf(configuration)
	if typ := configValue.Type(); typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("configuration should be a pointer to a struct type")
	}

	err = getFromYAML(filename, configuration)
	if err == nil {
		getFromEnvVariables(configuration)
	}

	return
}

func getFromYAML(filename string, configuration interface{}) (err error) {

	if len(filename) == 0 {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return
	}

	return
}

func getFromEnvVariables(configuration interface{}) {
	typ := reflect.TypeOf(configuration)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)

		// check if we've got a field name override for the environment
		tagContent := p.Tag.Get(envTagName)
		value := ""
		if len(tagContent) > 0 {
			value = os.Getenv(tagContent)
		} else {
			value = os.Getenv(p.Name)
		}

		if !p.Anonymous && len(value) > 0 {
			// struct
			s := reflect.ValueOf(configuration).Elem()

			if s.Kind() == reflect.Struct {
				// exported field
				f := s.FieldByName(p.Name)
				setValue(f, value)
			}
		}
	}
}

func setValue(f reflect.Value, value string) {
	// change value
	kind := f.Kind()

	if f.IsValid() && f.CanSet() {
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if kind == reflect.Int || kind == reflect.Int64 {
			setStringToInt(f, value, 64)
		} else if kind == reflect.Int32 {
			setStringToInt(f, value, 32)
		} else if kind == reflect.Int16 {
			setStringToInt(f, value, 16)
		} else if kind == reflect.Uint || kind == reflect.Uint64 {
			setStringToUInt(f, value, 64)
		} else if kind == reflect.Uint32 {
			setStringToUInt(f, value, 32)
		} else if kind == reflect.Uint16 {
			setStringToUInt(f, value, 16)
		} else if kind == reflect.Bool {
			setStringToBool(f, value)
		} else if kind == reflect.Float64 {
			setStringToFloat(f, value, 64)
		} else if kind == reflect.Float32 {
			setStringToFloat(f, value, 32)
		} else if kind == reflect.String {
			f.SetString(value)
		} else if kind == reflect.Struct {
			setJSONStringToStruct(f, value)
		} else if kind == reflect.Slice || kind == reflect.Array {
			setJSONStringToArray(f, value)
		}
	}

}

func setStringToInt(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseInt(value, 10, bitSize)

	if err == nil {
		if !f.OverflowInt(convertedValue) {
			f.SetInt(convertedValue)
		}
	}
}

func setStringToUInt(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseUint(value, 10, bitSize)

	if err == nil {
		if !f.OverflowUint(convertedValue) {
			f.SetUint(convertedValue)
		}
	}
}

func setStringToBool(f reflect.Value, value string) {
	convertedValue, err := strconv.ParseBool(value)

	if err == nil {
		f.SetBool(convertedValue)
	}
}

func setStringToFloat(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseFloat(value, bitSize)

	if err == nil {
		if !f.OverflowFloat(convertedValue) {
			f.SetFloat(convertedValue)
		}
	}
}

func setJSONStringToStruct(f reflect.Value, value string) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(value), &jsonMap)
	if err != nil {
		fmt.Errorf("Cannot decode string into map")
	}

	for k, v := range jsonMap {
		subField := f.Addr().Elem().FieldByName(k)
		setValue(subField, fmt.Sprintf("%v", v))
	}
}

func setJSONStringToArray(f reflect.Value, value string) {
	var jsonArr []interface{}
	err := json.Unmarshal([]byte(value), &jsonArr)
	if err != nil {
		fmt.Errorf("Cannot decode string into array")
		return
	}

	f.Set(reflect.MakeSlice(f.Type(), len(jsonArr), len(jsonArr)))
	for i, v := range jsonArr {
		jsonItemVal, err := json.Marshal(v)
		if err != nil {
			fmt.Errorf("Cannot marshall array item element")
			return
		}
		setValue(f.Index(i), string(jsonItemVal[:]))
	}
}
