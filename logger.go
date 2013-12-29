package logger

import (
  "fmt"
  "os"
  "runtime"
  "strings"
)

const (
  EMERG = iota
  CRIT
  ALERT
  ERR
  WARN
  NOTE
  INFO
  DEBUG
)

type Logger struct {
  Level int
  Debug bool
  Exit bool
}

func New() *Logger {
  return &Logger{
    Level: INFO,
    Debug: false,
    Exit: true,
  }
}

func (logger *Logger) Log(level int, text string) {
  if level > logger.Level {
    return
  }
  if logger.Debug != true && level == DEBUG {
    return
  }

  color := "90"
  label := "debug"

  if level == 6 {
    color = "32"
  }
  if level < 6 {
    color = "33"
  }
  if level < 4 {
    color = "31"
  }

  switch (level) {
    case INFO: label = "info"
    case NOTE: label = "notice"
    case WARN: label = "warning"
    case ERR: label = "error"
    case ALERT: label = "alert"
    case CRIT: label = "critical"
    case EMERG: label = "emergency"
  }

  color = "\x1b[" + color + "m"

  fmt.Printf("%s%s:%s %s\n", color, label, "\x1b[0m", text)

  if logger.Debug == true {
    _, file, line, ok := runtime.Caller(1)
    if ok == true {
      fmt.Printf("\x1b[90m%s%s, line %d\x1b[0m\n",
        strings.Repeat(" ", len(label) + 2), file, line)
    } else {
      fmt.Printf("\x1b[90m%sUnable to determine caller information.\x1b[0m\n",
        strings.Repeat(" ", len(label) + 2))
    }
  }

  if level < 2 && logger.Exit == true {
    os.Exit(1)
  }
}
