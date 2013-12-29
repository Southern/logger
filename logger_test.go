package logger

import "testing"

func TestLogger(t *testing.T) {
  logger := New()
  logger.Debug = true
  logger.Exit = false
  logger.Log(EMERG, "Test")
  logger.Log(CRIT, "Test")
  logger.Log(ALERT, "Test")
  logger.Log(ERR, "Test")
  logger.Log(WARN, "Test")
  logger.Log(NOTE, "Test")
  logger.Log(INFO, "Test")
  logger.Log(DEBUG, "Test")
}
