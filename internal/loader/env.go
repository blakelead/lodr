package loader

import (
	"fmt"
	"os"
	"reflect"
	"strings"

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
			unmarshalEnv(obj.Field(i), fmt.Sprintf("%s.%s", prefix, objType.Field(i).Name))
		}
	default:
		env := formatEnv(prefix)
		if val, ok := os.LookupEnv(env); ok {
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
				if prefix != "" {
					prefix = strings.ToUpper(prefix) + "_"
				}
				if val, ok := os.LookupEnv(prefix + name); ok {
					yaml.Unmarshal([]byte(val), obj.Field(i).Addr().Interface())
				}
			}
			unmarshalEnvFromTags(obj.Field(i), prefix)
		}
	}
}

func formatEnv(prefix string) string {
	return strings.ToUpper(strings.ReplaceAll(strings.Trim(prefix, "."), ".", "_"))
}
