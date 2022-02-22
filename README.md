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