# lodr

Minimal configuration loader for Go.

Load from file (yaml, json), from env or from command-line.

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

    // You can call File, Enc and Cmd in the order you like to change precedence at your preference.
    // You can also call File multiple times.
    Load(&c).
        File("config.yaml").
        Env().
        Cmd()

    fmt.Println(c.Name)
}
```
