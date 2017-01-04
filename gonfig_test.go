package gonfig

import (
	"os"
	"testing"
)

func GetFromJson_Filename_Empty_Should_Not_Panic(t *testing.T) {

	type Conf struct {
	}
	conf := Conf{}
	err := getFromJson("", &conf)

	if err != nil {
		t.Error("getFromJson should not panic", err)
	}
}

func getFromEnvVariables_should_not_panic_if_wrong_data(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_int(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_int16(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_int32(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_int64(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_uint(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_uint16(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_uint32(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_uint64(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_bool(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_float32(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_float64(t *testing.T) {
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

func getFromEnvVariables_should_find_and_parse_string(t *testing.T) {
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
