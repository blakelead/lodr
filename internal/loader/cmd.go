package loader

import (
	"flag"
	"reflect"
	"time"
)

// LoadCmd defines flags from cmd tags then
// inject parsed flags values in out object
func LoadCmd(obj interface{}) {
	objValue := reflect.ValueOf(obj)
	cmds := make(map[string]interface{})
	createFlags(objValue, cmds)
	flag.Parse()
	unmarshalCmd(objValue, cmds)
}

func createFlags(obj reflect.Value, cmds map[string]interface{}) {
	switch obj.Kind() {
	case reflect.Ptr, reflect.Interface:
		createFlags(obj.Elem(), cmds)
	case reflect.Struct:
		objType := obj.Type()
		for i := 0; i < objType.NumField(); i++ {
			if name, ok := obj.Type().Field(i).Tag.Lookup("cmd"); ok {
				switch obj.Field(i).Kind() {
				case reflect.String:
					cmds[name] = flag.String(name, obj.Field(i).String(), "")
				case reflect.Int:
					cmds[name] = flag.Int(name, int(obj.Field(i).Int()), "")
				case reflect.Float64:
					cmds[name] = flag.Float64(name, obj.Field(i).Float(), "")
				case reflect.Bool:
					cmds[name] = flag.Bool(name, obj.Field(i).Bool(), "")
				case reflect.TypeOf(time.Second).Kind():
					val := obj.Field(i).Interface().(time.Duration)
					cmds[name] = flag.Duration(name, val, "")
				}
			}
			createFlags(obj.Field(i), cmds)
		}
	}
}

func unmarshalCmd(obj reflect.Value, cmds map[string]interface{}) {
	switch obj.Kind() {
	case reflect.Ptr, reflect.Interface:
		unmarshalCmd(obj.Elem(), cmds)
	case reflect.Struct:
		objType := obj.Type()
		for i := 0; i < objType.NumField(); i++ {
			if name, ok := obj.Type().Field(i).Tag.Lookup("cmd"); ok {
				if _, ok := cmds[name]; !ok {
					continue
				}
				valPtr := reflect.ValueOf(cmds[name])
				val := reflect.Indirect(valPtr)
				obj.Field(i).Set(val)
			}
			unmarshalCmd(obj.Field(i), cmds)
		}
	}
}
