// Package lodr is a simple configuration file loader that
// loads JSON and YAML file formats, environment variables
// and command-line flags.
//
// Example:
//
// ```yaml
// # config.yaml
// name: mydb
// db:
//   host: localhost
//   port: 8080
//   timeout: 10s
//   tls: true
// ```
//
// ```go
// package main
//
// import (
//     "fmt"
//     "time"
//     "github.com/blakelead/lodr"
// )
//
// type MyConfig struct {
//     Name string `cmd:"name"`
//     DB   struct {
//         Host     string        `yaml:"host" env:"DB_HOST" cmd:"db.host"`
//         Port     int           `yaml:"port" env:"DB_PORT" cmd:"db.port"`
//         Password string        `env:"DB_PASSWORD"`
//         Timeout  time.Duration `yaml:"timeout" env:"DB_TIMEOUT" cmd:"db.timeout"`
//         TLS      bool          `yaml:"tls" env:"DB_TLS" cmd:"db.tls"`
//     }
// }
//
// func main() {
//     var mc MyConfig
//
//     opts := &lodr.EnvOptions{
//         Prefix: "MY_APP",
//     }
//
//     c := lodr.Load(&mc).File("config.yaml").EnvWithOptions(opts).Cmd()
//
//     if c.Error != nil {
//         panic(c.Error)
//     }
//
//     fmt.Println(mc.Name)
//     fmt.Println(mc.DB.Password)
//     fmt.Printf("%s:%d\n", mc.DB.Host, mc.DB.Port)
// }
// ```
//
// ```bash
// > MY_APP_DB_PASSWORD=pass go run main.go --name the_db
// the_db
// pass
// localhost:8080
// ```
//
// Go to https://github.com/blakelead/lodr for more examples.
package lodr

import (
	"github.com/blakelead/lodr/internal/loader"
)

// Config structure.
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
	if c.Error == nil {
		if err := loader.LoadFile(filename, c.object); err != nil {
			c.Error = err
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
	if c.Error == nil {
		if err := loader.LoadEnv(c.object, &loader.EnvOptions{
			Prefix:     opts.Prefix,
			ProcessAll: opts.ProcessAll,
		}); err != nil {
			c.Error = err
		}
	}
	return c
}

// Cmd parses command-line flags and loads them in user object.
func (c *Config) Cmd() *Config {
	if c.Error == nil {
		loader.LoadCmd(c.object)
	}
	return c
}
