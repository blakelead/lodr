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

// EnvOptions is a helper object
// to pass options to the loader
type EnvOptions struct {
	Prefix     string
	ProcessAll bool
}

// Load func
func Load(in interface{}) *Config {
	c := &Config{
		object: in,
	}
	return c
}

// File func
func (c *Config) File(filename string) *Config {
	err := loader.LoadFile(filename, &c.object)
	if err != nil {
		c.Error = fmt.Errorf("%w\n%s", c.Error, err.Error())
	}
	return c
}

// Env func
func (c *Config) Env(opts *EnvOptions) *Config {

	err := loader.LoadEnv(&c.object, &loader.EnvOptions{
		Prefix:     opts.Prefix,
		ProcessAll: opts.ProcessAll,
	})
	if err != nil {
		c.Error = fmt.Errorf("%w\n%s", c.Error, err.Error())
	}
	return c
}

// Cmd func
func (c *Config) Cmd() *Config {
	loader.LoadCmd(&c.object)
	return c
}
