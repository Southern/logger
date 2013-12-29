package logger

import (
  "fmt"
  "os"
  "regexp"
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

var RegexMap = map[*regexp.Regexp]int {
  regexp.MustCompile("(?i)^e(mer(gency)?)?$"): EMERG,
  regexp.MustCompile("(?i)^c(rit(ical)?)?$"): CRIT,
  regexp.MustCompile("(?i)^a(lert)?$"): ALERT,
  regexp.MustCompile("(?i)^e(rr(or)?)?$"): ERR,
  regexp.MustCompile("(?i)^w(arn(ing)?)?$"): WARN,
  regexp.MustCompile("(?i)^n(ot(e|ice)?)?$"): NOTE,
  regexp.MustCompile("(?i)^i(nfo(rmation)?)?$"): INFO,
  regexp.MustCompile("(?i)^d(ebug)?$"): DEBUG,
}

type Logger struct {
  Level string
  Debug bool
  Exit bool
}

func New() *Logger {
  return &Logger{
    Level: "i",
    Debug: false,
    Exit: true,
  }
}

func (logger *Logger) GetLevel(level string) (int) {
  for regexp, i := range RegexMap {
    if regexp.MatchString(level) {
      return i
    }
  }
  return -1
}

func (logger *Logger) Log(data ...string) (*Logger) {
  level := logger.GetLevel(data[0])
  if level > -1 {
    data = data[1:]
    return logger.Raw(level, data...)
  }

  return logger.Raw(INFO, data...)
}

func (logger *Logger) Raw(level int, text ...string) (*Logger) {
  if level > logger.GetLevel(logger.Level) {
    return logger
  }
  if logger.Debug != true && level == DEBUG {
    return logger
  }

  if text[0] == string(level) {
    text = text[1:]
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

  for i := 0; i < len(text); i++ {
    fmt.Printf("%s%s:%s %s\n", color, label, "\x1b[0m", text[i])

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
  }

  if level < 2 && logger.Exit == true {
    os.Exit(1)
  }

  return logger
}
