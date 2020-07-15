# lodr

Minimal configuration loader for Go.

Load from file (YAML, JSON), from environment variables or from command-line flags.

## Usage

### Load from file

```yaml
# config.yaml
name: value
```

```go
package main

import (
    "fmt"
    "github.com/blakelead/lodr"
)

type MyConfig struct {
    Name string
}

func main() {
    var myConfig MyConfig

    lodr.Load(&myConfig).File("config.yaml")

    fmt.Println(myConfig.Name)
}
```

```bash
> go run main.go
value
```

### Load from file with tag

It uses `gopkg.in/yaml.v2` so it works exactly the same:

```yaml
# config.yaml
my_app_name: value
```

```go
type MyConfig struct {
    Name string `yaml:"my_app_name"`
}

func main() {
    var myConfig MyConfig

    lodr.Load(&myConfig).File("config.yaml")

    fmt.Println(myConfig.Name)
}
```

```bash
> go run main.go
value
```

### Load from env

```go
type MyConfig struct {
    Name string
}

func main() {

    var myConfig MyConfig

    lodr.Load(&myConfig).Env()

    fmt.Println(myConfig.Name)
}
```

```bash
> NAME=new_value go run main.go
new_value
```

### Load from env with tag

```go
type MyConfig struct {
    Name string `env:MY_APP_NAME`
}

func main() {

    var myConfig MyConfig

    lodr.Load(&myConfig).Env()

    fmt.Println(myConfig.Name)
}
```

```bash
> MY_APP_NAME=value go run main.go
value
```

### Load from env with options

```go
type MyConfig struct {
    Name string `env:NAME`
}

func main() {

    var myConfig MyConfig

    opts := &lodr.EnvOptions{
          Prefix:     "MY_APP",
          ProcessAll: false,
    }

    lodr.Load(&myConfig).EnvWithOptions(opts)

    fmt.Println(myConfig.Name)
}
```

```bash
> MY_APP_NAME=value go run main.go
value
```

Options:

- `Prefix`: environment variables are all prefixed by this value
- `ProcessAll`: if true, no need to specify tags. Names will be infered from attributes.

### Load from command-line arguments

 main.go

 ```go
type MyConfig struct {
    Name string `cmd:myapp.name`
}

func main() {

    var myConfig MyConfig

    lodr.Load(&myConfig).Cmd()

    fmt.Println(myConfig.Name)
}
 ```

```bash
> go run main.go --myapp.name the_value
the_value
```

### Complete example

```yaml
# config.yaml
name: mydb
db:
  host: localhost
  port: 8080
  timeout: 10s
  tls: true
```

```go
package main

import (
    "fmt"
    "time"
    "github.com/blakelead/lodr"
)

type MyConfig struct {
    Name string `cmd:"name"`
    DB   struct {
        Host     string        `yaml:"host" env:"DB_HOST" cmd:"db.host"`
        Port     int           `yaml:"port" env:"DB_PORT" cmd:"db.port"`
        Password string        `env:"DB_PASSWORD"`
        Timeout  time.Duration `yaml:"timeout" env:"DB_TIMEOUT" cmd:"db.timeout"`
        TLS      bool          `yaml:"tls" env:"DB_TLS" cmd:"db.tls"`
    }
}

func main() {
    var mc MyConfig

    opts := &lodr.EnvOptions{
        Prefix: "MY_APP",
    }

    c := lodr.Load(&mc).File("config.yaml").EnvWithOptions(opts).Cmd()

    if c.Error != nil {
        panic(c.Error)
    }

    fmt.Println(mc.Name)
    fmt.Println(mc.DB.Password)
    fmt.Printf("%s:%d\n", mc.DB.Host, mc.DB.Port)
}
```

```bash
> MY_APP_DB_PASSWORD=pass go run main.go --name the_db
the_db
pass
localhost:8080
```
