package lodr

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestJsonFile(t *testing.T) {
	generateTestFiles(t)
	defer deleteTestFiles(t)

	type testConfig struct {
		StringParam string        `json:"jsonStringParam"`
		IntParam    int           `json:"jsonIntParam"`
		FloatParam  float64       `json:"jsonFloatParam"`
		TimeParam   time.Duration `json:"jsonTimeParam"`
		BoolParam   bool          `json:"jsonBoolParam"`
	}
	var tc testConfig

	Load(&tc).File("config.json")

	if tc.StringParam != "json_string_param" {
		t.Errorf("Should be json_string_param, is %s", tc.StringParam)
	}
	if tc.IntParam != 3 {
		t.Errorf("Should be 3, is %d", tc.IntParam)
	}
	if tc.FloatParam != 3.1415 {
		t.Errorf("Should be 3.1415, is %f", tc.FloatParam)
	}
	if tc.TimeParam != time.Second*15 {
		t.Errorf("Should be 15s, is %s", tc.TimeParam)
	}
	if tc.BoolParam != true {
		t.Errorf("Should be true, is %t", tc.BoolParam)
	}
}

func TestEnvWithPrefix(t *testing.T) {
	type testConfig struct {
		StringParam string        `env:"STRING_PARAM"`
		IntParam    int           `env:"INT_PARAM"`
		FloatParam  float64       `env:"FLOAT_PARAM"`
		TimeParam   time.Duration `env:"TIME_PARAM"`
		BoolParam   bool          `env:"BOOL_PARAM"`
		Nested      struct {
			StringParam string `env:"NESTED_STRING_PARAM"`
		} `env:"NESTED_STRUCT"`
	}
	var tc testConfig

	os.Setenv("TEST_ENV_STRING_PARAM", "env_string_param")
	os.Setenv("TEST_ENV_INT_PARAM", "4")
	os.Setenv("TEST_ENV_FLOAT_PARAM", "4.1234")
	os.Setenv("TEST_ENV_TIME_PARAM", "10s")
	os.Setenv("TEST_ENV_BOOL_PARAM", "true")
	os.Setenv("TEST_ENV_NESTED_STRING_PARAM", "env_nested_string_param")

	opts := &EnvOptions{
		Prefix:     "TEST_ENV",
		ProcessAll: false,
	}

	Load(&tc).EnvWithOptions(opts)
	if tc.StringParam != "env_string_param" {
		t.Errorf("Should be env_string_param, is %s", tc.StringParam)
	}
	if tc.IntParam != 4 {
		t.Errorf("Should be 4, is %d", tc.IntParam)
	}
	if tc.FloatParam != 4.1234 {
		t.Errorf("Should be 4.1234, is %f", tc.FloatParam)
	}
	if tc.TimeParam != time.Second*10 {
		t.Errorf("Should be 10s, is %s", tc.TimeParam)
	}
	if tc.BoolParam != true {
		t.Errorf("Should be true, is %t", tc.BoolParam)
	}
	if tc.Nested.StringParam != "env_nested_string_param" {
		t.Errorf("Should be env_nested_string_param, is %s", tc.Nested.StringParam)
	}
}

func TestEnvWithProcessAll(t *testing.T) {
	type testConfig struct {
		StringParam string        `env:"STRINGPARAM"`
		IntParam    int           `env:"INTPARAM"`
		FloatParam  float64       `env:"FLOATPARAM"`
		TimeParam   time.Duration `env:"TIMEPARAM"`
		BoolParam   bool          `env:"BOOLPARAM"`
	}
	var tc testConfig

	os.Setenv("STRING_PARAM", "env_string_param")
	os.Setenv("INT_PARAM", "4")
	os.Setenv("FLOAT_PARAM", "4.1234")
	os.Setenv("TIME_PARAM", "10s")
	os.Setenv("BOOL_PARAM", "true")

	opts := &EnvOptions{
		ProcessAll: true,
	}

	Load(&tc).EnvWithOptions(opts)
	if tc.StringParam != "env_string_param" {
		t.Errorf("Should be env_string_param, is %s", tc.StringParam)
	}
	if tc.IntParam != 4 {
		t.Errorf("Should be 4, is %d", tc.IntParam)
	}
	if tc.FloatParam != 4.1234 {
		t.Errorf("Should be 4.1234, is %f", tc.FloatParam)
	}
	if tc.TimeParam != time.Second*10 {
		t.Errorf("Should be 10s, is %s", tc.TimeParam)
	}
	if tc.BoolParam != true {
		t.Errorf("Should be true, is %t", tc.BoolParam)
	}
}

func generateTestFiles(t *testing.T) {
	configJSON := []byte(`{
	"jsonStringParam": "json_string_param",
	"jsonIntParam": 3,
	"jsonFloatParam": 3.1415,
	"jsonTimeParam": 15000000000,
	"jsonBoolParam": true
}`)
	err := ioutil.WriteFile("config.json", configJSON, 0644)
	if err != nil {
		t.Error("Could not generate test file config.json")
	}
}

func deleteTestFiles(t *testing.T) {
	files := []string{
		"config.json",
	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			t.Errorf("Could not delete test file %s", file)
		}
	}
}
