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

type LoggerStack struct {
  Errors int
  Debug int
}

type Logger struct {
  Debug bool
  Colorize bool
  Exit bool
  Level string
  Stack int
}

func stack(entries int) (stack map[int]map[string]interface{}) {
  stack = make(map[int]map[string]interface{})

  for i := 2; i < entries + 2; i++ {
    _, file, line, ok := runtime.Caller(i)
    if !ok {
      return
    }

    id := len(stack) + 1
    stack[id] = make(map[string]interface{})
    stack[id]["file"] = file
    stack[id]["line"] = line
  }
  return
}

func New() *Logger {
  return &Logger{
    Debug: false,
    Colorize: true,
    Exit: true,
    Level: "i",
    Stack: 25,
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

  color := 90
  label := "debug"

  if level == 6 {
    color = 32
  }
  if level < 6 {
    color = 33
  }
  if level < 4 {
    color = 31
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

  label = label + ":"

  formats := map[string]string {
    "label": fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, label),
    "debug": fmt.Sprintf("%s%%s, line %%d",
      strings.Repeat(" ", len(label) + 1)),
    "debugColor": "\x1b[90m%s\x1b[0m",
  }

  for i := 0; i < len(text); i++ {
    if logger.Colorize {
      fmt.Printf("%s %s\n", formats["label"], text[i])
    } else {
      fmt.Printf("%s %s\n", label, text[i])      
    }

    if logger.Debug == true || level <= ERR {
      callers := stack(logger.Stack)

      for x := 1; x < len(callers) + 1; x++ {
        stack := callers[x]
        if logger.Colorize {
          fmt.Printf(formats["debugColor"] + "\n",
            fmt.Sprintf(formats["debug"], stack["file"], stack["line"]))
        } else {
          fmt.Printf(formats["debug"] + "\n", stack["file"], stack["line"])
        }
      }
    }
  }

  if level < 2 && logger.Exit == true {
    os.Exit(1)
  }

  return logger
}
