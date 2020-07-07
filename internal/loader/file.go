package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// LoadFile unmarshals yaml or json file into obj
func LoadFile(filename string, obj interface{}) error {
	blob, err := read(filename)
	if err != nil {
		return err
	}
	return unmarshal(filename, blob, obj)
}

// readfile checks  if the provided filename is a  valid path to
// the file. If it is not, it checks if the filename corresponds
// to a file relative to the executable directory. It then reads
// the file and returns its content.
func read(filename string) ([]byte, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		absPath, err := os.Executable()
		if err != nil {
			return []byte{}, err
		}
		filename = filepath.Dir(absPath) + string(os.PathSeparator) + filename
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return []byte{}, err
		}
	}
	if fi, _ := os.Stat(filename); fi.Size() == 0 {
		return []byte{}, errors.New("Configuration file is empty")
	}
	return ioutil.ReadFile(filename)
}

// unmarshal calls yaml or json package Unmarshal depending on the
// filename extension.
func unmarshal(filename string, blob []byte, obj interface{}) error {
	switch ext := filepath.Ext(filename); ext {
	case ".yaml", ".yml":
		err := yaml.Unmarshal(blob, obj)
		if err != nil {
			return err
		}
	case ".json":
		err := json.Unmarshal(blob, obj)
		if err != nil {
			return err
		}
	case "":
		return fmt.Errorf("Invalid file name: extension missing")
	default:
		return fmt.Errorf("File type '%s' is not supported", ext)
	}
	return nil
}
