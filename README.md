# [NathanBurkett](https://github.com/NathanBurkett) / [env](https://github.com/NathanBurkett/env)
[![Documentation](https://godoc.org/github.com/nathanburkett/env?status.svg)](http://godoc.org/github.com/nathanburkett/env) [![Coverage Status](https://coveralls.io/repos/github/NathanBurkett/env/badge.svg?branch=master)](https://coveralls.io/github/NathanBurkett/env?branch=master) [![CircleCI](https://circleci.com/gh/NathanBurkett/env.svg?style=svg)](https://circleci.com/gh/NathanBurkett/env) [![Go Report Card](https://goreportcard.com/badge/github.com/nathanburkett/env)](https://goreportcard.com/report/github.com/nathanburkett/env)

View the [documentation](https://godoc.org/github.com/nathanburkett/env)

## Install
```bash
go get -u github.com/nathanburkett/env
```

## Usage
### Reading an environmental file via [`Reader`](https://godoc.org/github.com/NathanBurkett/env#Reader)
```go
package main

import (
    "fmt"
    "log"
    "github.com/nathanburkett/env"
    "os"
)

func main() {
    pwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    readEnv(pwd)
}

func readEnv(workingDir string) {
    envPath := fmt.Sprintf("%s/.env", workingDir)

    file, err := os.Open(envPath)
    if err != nil {
        log.Panic(err)
	}
    defer file.Close()

    env.NewReader(file).Read()
}
```

### Using [`Must()`](https://godoc.org/github.com/NathanBurkett/env#Must)
```go
package main

import (
    "github.com/jmoiron/sqlx"
    "github.com/nathanburkett/env"

    "log"
)

func main() {
    _, err := sqlx.Open("mysql", env.Must("DB_DSN"))
    if err != nil {
        log.Fatal(err)
    }

    // ...
}
```

## Running tests
On MacOS you can install or upgrade to the latest released version with Homebrew:

```bash
brew install dep
brew upgrade dep
```

On other platforms you can use the install.sh script:

```bash
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

```bash
make install -B
make coverage -B
```

## License
Copyright (c) 2018 Nathan Burkett

Licensed under the [MIT License](LICENSE)
