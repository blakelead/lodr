package loader

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/blakelead/lodr/internal/utils"
	"gopkg.in/yaml.v2"
)

// EnvOptions is a helper object
// to pass options to the loader
type EnvOptions struct {
	Prefix     string
	ProcessAll bool
}

// LoadEnv set obj from environment variables
func LoadEnv(obj interface{}, opts *EnvOptions) error {
	objValue := reflect.ValueOf(obj)
	if opts.ProcessAll {
		unmarshalEnv(objValue, opts.Prefix)
	} else {
		unmarshalEnvFromTags(objValue, opts.Prefix)
	}
	return nil
}

func unmarshalEnv(obj reflect.Value, prefix string) {
	switch obj.Kind() {
	case reflect.Ptr, reflect.Interface:
		unmarshalEnv(obj.Elem(), prefix)
	case reflect.Struct:
		objType := obj.Type()
		for i := 0; i < objType.NumField(); i++ {
			fieldName := utils.SplitCamelCase(objType.Field(i).Name, "_")
			unmarshalEnv(obj.Field(i), fmt.Sprintf("%s.%s", prefix, fieldName))
		}
	default:
		envVar := strings.ReplaceAll(strings.Trim(prefix, "."), ".", "_")
		fmt.Println(envVar)
		if val, ok := os.LookupEnv(strings.ToUpper(envVar)); ok {
			yaml.Unmarshal([]byte(val), obj.Addr().Interface())
		}
	}
}

func unmarshalEnvFromTags(obj reflect.Value, prefix string) {
	switch obj.Kind() {
	case reflect.Ptr, reflect.Interface:
		unmarshalEnvFromTags(obj.Elem(), prefix)
	case reflect.Struct:
		objType := obj.Type()
		for i := 0; i < objType.NumField(); i++ {
			if name, ok := obj.Type().Field(i).Tag.Lookup("env"); ok {
				envVar := name
				if prefix != "" {
					envVar = prefix + "_" + envVar
				}
				if val, ok := os.LookupEnv(strings.ToUpper(envVar)); ok {
					yaml.Unmarshal([]byte(val), obj.Field(i).Addr().Interface())
				}
			}
			unmarshalEnvFromTags(obj.Field(i), prefix)
		}
	}
}
