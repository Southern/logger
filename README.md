# Go logger

## Installation
```
go get github.com/southern/logger
```

## Testing
```
go test github.com/southern/logger
```

## Usage

### Logger
```go
type Logger struct {
  Level int
  Debug bool
  Exit bool
}
```

### Log levels
```go
const (
  // Emergency
  EMERG = iota

  // Critical
  CRIT

  // Alert
  ALERT

  // Error
  ERR

  // Warning
  WARN

  // Notice
  NOTE

  // Info
  INFO

  // Debug
  DEBUG
)
```

### Example
```go
package main

import "github.com/southern/logger"

func main() {
  // Initialize a new log with the default values. &logger.Logger{} is another
  // way it can be constructed if you don't want to call .New()
  log := logger.New()

  // Messages must be this level or a more serious level in order for it to be
  // output
  log.Level = logger.Info

  // Allow debug messages to be output. This includes a line showing where the
  // log call originated.
  log.Debug = true

  // Automatically exit on logger.EMERG or logger.CRIT
  log.Exit = false

  log.Log(EMERG, "This is an emergency.")
  log.Log(CRIT, "This is a critical message.")
  log.Log(ALERT, "This is an alert.")
  log.Log(ERR, "This is an error.")
  log.Log(WARN, "This is a warning.")
  log.Log(NOTE, "This is a notice.")
  log.Log(INFO, "This is an informational message.")
  log.Log(DEBUG, "This is a debug message.")
}
```

## License
Copyright (c) 2013 Colton Baker

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
