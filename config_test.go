package gonfig_test

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	gonfig "github.com/g41797/gonfig"
)

type configuration struct {
	Port              int
	Connection_String string
}

//
// because mixed testing environment, these test should be activated manually
//

func Test_unmarshalFromJSONFile_ToStruct(t *testing.T) {
	expected := defaults()
	actual := configuration{}
	fPath := "./_config/example.json"
	err := gonfig.GetConf(fPath, &actual)
	compare(t, err, actual, expected)
}

func Test_unmarshalFromJSONFile_AndEnv_ToStruct(t *testing.T) {
	expected := defaults()
	expected.Port = 8443

	os.Setenv("Port", strconv.Itoa((expected.Port)))

	actual := configuration{}
	fPath := "./_config/example.json"
	err := gonfig.GetConf(fPath, &actual)
	compare(t, err, actual, expected)
}

// ------------------------------------------------------------------------------------------
// -- Manually run test under vscode:
// go test -run ^Test_unmarshalFromJSONFile_AndPrefixedEnv_ToStruct$ github.com/tkanos/gonfig
// ------------------------------------------------------------------------------------------
func Test_unmarshalFromJSONFile_AndPrefixedEnv_ToStruct(t *testing.T) {
	expected := defaults()
	expected.Port = 8443
	fPath := "./_config/example.json"

	os.Setenv(envPrefix(fPath)+"Port", strconv.Itoa((expected.Port)))

	actual := configuration{}
	err := gonfig.GetConf(fPath, &actual)
	compare(t, err, actual, expected)
}

func defaults() configuration {
	return configuration{Port: 8080, Connection_String: "connection_string"}
}

func compare(t *testing.T, getErr error, actual configuration, expected configuration) {
	if getErr != nil {
		t.Errorf("unmarshal error %v", getErr)
	}

	if actual.Connection_String != expected.Connection_String {
		t.Errorf("Expected %s Actual %s", expected.Connection_String, actual.Connection_String)
	}

	if actual.Port != expected.Port {
		t.Errorf("Expected %d Actual %d", expected.Port, actual.Port)
	}
}

func envPrefix(fPath string) string {
	prefix := strings.ToUpper(baseFileName(fPath)) + "_"
	return prefix
}

func baseFileName(fPath string) string {
	if len(fPath) == 0 {
		return ""
	}

	file := filepath.Base(fPath)
	ext := filepath.Ext(file)
	name := strings.TrimSuffix(file, ext)

	return name
}
