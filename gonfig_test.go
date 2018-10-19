package gonfig

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_GetFromYAML_Filename_Empty_Should_Not_Panic(t *testing.T) {

	type Conf struct {
	}
	conf := Conf{}
	err := getFromYAML("", &conf)

	if err != nil {
		t.Error("getFromYAML should not panic", err)
	}
}

func tmpFileWithContent(content string, t *testing.T) string {

	file, err := ioutil.TempFile("", "gonfig_test_data_")
	if err != nil {
		t.Error("Error creating file with test data", err)
	}

	_, err = io.Copy(file, strings.NewReader("{}"))
	if err != nil {
		t.Error("Error writing test data", err)
	}

	return file.Name()
}

func Test_GetFromYAML_Filename_Should_Not_be_Panic(t *testing.T) {

	filename := tmpFileWithContent("{}", t)
	defer os.Remove(filename)

	type Conf struct {
	}
	conf := Conf{}
	err := getFromYAML(filename, &conf)

	if err != nil {
		t.Error("getFromYAML file not found", err)
	}
}

func Test_getFromEnvVariables_should_not_panic_if_wrong_data(t *testing.T) {
	type Conf struct {
		Id int
	}
	os.Setenv("Id", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 0 {
		t.Error("Id should be 0", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int(t *testing.T) {
	type Conf struct {
		Id int
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int16(t *testing.T) {
	type Conf struct {
		Id int16
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int32(t *testing.T) {
	type Conf struct {
		Id int32
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int64(t *testing.T) {
	type Conf struct {
		Id int64
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint(t *testing.T) {
	type Conf struct {
		Id uint
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint16(t *testing.T) {
	type Conf struct {
		Id uint16
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint32(t *testing.T) {
	type Conf struct {
		Id uint32
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint64(t *testing.T) {
	type Conf struct {
		Id uint64
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_bool(t *testing.T) {
	type Conf struct {
		Id bool
	}
	os.Setenv("Id", "true")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != true {
		t.Error("Id should be true", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_float32(t *testing.T) {
	type Conf struct {
		Id float32
	}
	os.Setenv("Id", "123.123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123.123 {
		t.Error("Id should be 123.123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_float64(t *testing.T) {
	type Conf struct {
		Id float64
	}
	os.Setenv("Id", "123.123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123.123 {
		t.Error("Id should be 123.123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		Id string
	}
	os.Setenv("Id", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != "abc" {
		t.Error("Id should be abc", conf.Id)
	}
}

func Test_getFromCustomEnvVariables_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		Id string `env:"CONF_ID"`
	}
	os.Setenv("CONF_ID", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != "abc" {
		t.Error("Id should be abc", conf.Id)
	}
}
