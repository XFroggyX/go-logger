## go-logger

A small wrapper for logging

## Install

```bash
$ go get github.com/XFroggyX/go-logger
```


## Getting Started
Getting a handle and creating a folder with a file

```go

import (
    lgr "github.com/XFroggyX/go-logger"
)

folder := "logs"
fileLogName := "project_log.log"
l := lgr.Init(folder, fileLogName)
```

## Usage

```go
_, err := strconv.Atoi("123ABC")
l.Logf("INFO some important message, %v", err)
```

####output looks like this:

stdout:
```text
2022/02/10 12:32:47 INFO  some important message, strconv.Atoi: parsing "123ABC": invalid syntax
```

file log:
```text
2022/02/10 12:32:47.315 INFO  (go-logger@v0.0.0-20220210091434-4af4ac24af25/logging.go:16 go-logger.(*Logger).Logf) some important message, strconv.Atoi: parsing "123ABC": invalid syntax
```