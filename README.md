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
  // output.
  log.Level = "debug"

  // Allow debug information to be output. This includes a line showing where
  // the log call originated. This DOES NOT set log.Level to "debug" to show
  // debug messages.
  log.Debug = true

  // Automatically exit on logger.EMERG or logger.CRIT
  log.Exit = false

  logger.Log("e", "This is an emergency.")
  logger.Log("emer", "This is an emergency.")
  logger.Log("emergency", "This is an emergency.")

  logger.Log("c", "This is a critical message.")
  logger.Log("crit", "This is a critical message.")
  logger.Log("critical", "This is a critical message.")

  logger.Log("a", "This is an alert.")
  logger.Log("alert", "This is an alert.")

  logger.Log("e", "This is an error.")
  logger.Log("err", "This is an error.")
  logger.Log("error", "This is an error.")

  logger.Log("w", "This is a warning.")
  logger.Log("warn", "This is a warning.")
  logger.Log("warning", "This is a warning.")

  logger.Log("n", "This is a notice.")
  logger.Log("note", "This is a notice.")
  logger.Log("notice", "This is a notice.")

  logger.Log("This is an informational message.")
  logger.Log("i", "This is an informational message.")
  logger.Log("info", "This is an informational message.")
  logger.Log("information", "This is an informational message.")

  logger.Log("d", "This is a debug message.")
  logger.Log("debug", "This is a debug message.")

  // You can also use multiple messages in one Log function.
  logger.Log("This is an informational message.",
    "This is another informational message")
}
```

## License
Copyright (c) 2013 Colton Baker

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
