# lodr

Minimal configuration loader for Go.

Load from file (YAML, JSON), from environment variables or from command-line flags.

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/blakelead/lodr"
)

// MyConfig correspond to your configuration struct.
type MyConfig struct {
    Name string `yaml:"name" env:"MY_APP_NAME" cmd:"app.name"`
    DB struct {
        Host string `yaml:"host" env:"MY_APP_DB_HOST" cmd:"app.db.host"`
        Port int `yaml:"port" env:"MY_APP_DB_PORT" cmd:"app.db.port"`
        Timeout time.Duration `yaml:"timeout" env:"MY_APP_DB_TIMEOUT" cmd:"app.db.timeout"`
    }
}

func main() {

    var c MyConfig

    // Options can be passed to Env(). See details below.
    opts := &loader.EnvOptions{
        Prefix:     "TEST_ENV",
        ProcessAll: false,
    }

    // You can call File(), Env() and Cmd() in the order you need. The order determines precedence.
    Load(&c).
        File("config.yaml").
        Env(opts).
        Cmd()

    fmt.Println(c.Name)
}
```

## Files

YAML and JSON formats are supported for now. You can call File() multiple times, in the order you want.
Files are the only way to specify complex types such as arrays and maps.

## Env

You can use Env() to define values for simple types (string, int, float, bool, time.Duration) in the root of the struct and in nested structs.

Environment variables will always be looked up in uppercase.

The following options can be passed:

- `Prefix`: you can specify a prefix to avoid writing the same prefix in the environment variables you define
- `ProcessAll`: if true, you don't need to define environment variables tags. The names and structure of your object will be used to infer the names.

## Cmd

You can use Cmd() to define values for simple types (string, int, float, bool, time.Duration) in the root of the struct and in nested structs.
