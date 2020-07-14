package lodr

import (
	"fmt"

	"github.com/blakelead/lodr/internal/loader"
)

// Config struct
type Config struct {
	object interface{}
	errors []string
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
		c.errors = append(c.errors, err.Error())
	}
	return c
}

// Env func
func (c *Config) Env() *Config {
	err := loader.LoadEnv(&c.object, &loader.EnvOptions{})
	if err != nil {
		c.errors = append(c.errors, err.Error())
	}
	return c
}

// Cmd func
func (c *Config) Cmd() *Config {
	loader.LoadCmd(&c.object)
	return c
}

// Run func
func (c *Config) Run() error {
	if len(c.errors) != 0 {
		return fmt.Errorf("%v", c.errors)
	}
	return nil
}
