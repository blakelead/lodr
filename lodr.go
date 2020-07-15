package lodr

import (
	"fmt"

	"github.com/blakelead/lodr/internal/loader"
)

// Config struct
type Config struct {
	object interface{}
	Error  error
}

// EnvOptions is a helper object to pass options to the loader.
type EnvOptions struct {
	Prefix     string
	ProcessAll bool
}

// Load stores the user given object in Config.
func Load(in interface{}) *Config {
	return &Config{
		object: in,
	}
}

// File unmarshal file content into the user object.
func (c *Config) File(filename string) *Config {
	err := loader.LoadFile(filename, c.object)
	if err != nil {
		if c.Error == nil {
			c.Error = err
		} else {
			c.Error = fmt.Errorf("%w\n%s", c.Error, err.Error())
		}
	}
	return c
}

// Env reads environment variables and loads them in user object.
func (c *Config) Env() *Config {
	return c.EnvWithOptions(&EnvOptions{})
}

// EnvWithOptions is like Env() but allows to pass options.
func (c *Config) EnvWithOptions(opts *EnvOptions) *Config {
	err := loader.LoadEnv(c.object, &loader.EnvOptions{
		Prefix:     opts.Prefix,
		ProcessAll: opts.ProcessAll,
	})
	if err != nil {
		if c.Error == nil {
			c.Error = err
		} else {
			c.Error = fmt.Errorf("%w\n%s", c.Error, err.Error())
		}
	}
	return c
}

// Cmd func
func (c *Config) Cmd() *Config {
	loader.LoadCmd(c.object)
	return c
}
