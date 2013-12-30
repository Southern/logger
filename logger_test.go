package logger

import "testing"

func TestLogger(t *testing.T) {
  logger := New()
  logger.Exit = false

  // Test Raw
  logger.Raw(EMERG, "Test")
  logger.Raw(CRIT, "Test")
  logger.Raw(ALERT, "Test")
  logger.Raw(ERR, "Test")
  logger.Raw(WARN, "Test")
  logger.Raw(NOTE, "Test")
  logger.Raw(INFO, "Test")
  logger.Raw(DEBUG, "Test")

  logger.Log("e", "Test")
  logger.Log("emer", "Test")
  logger.Log("emergency", "Test")

  logger.Log("c", "Test")
  logger.Log("crit", "Test")
  logger.Log("critical", "Test")

  logger.Log("a", "Test")
  logger.Log("alert", "Test")

  logger.Log("err", "Test")
  logger.Log("error", "Test")

  logger.Log("w", "Test")
  logger.Log("warn", "Test")
  logger.Log("warning", "Test")

  logger.Log("n", "Test")
  logger.Log("note", "Test")
  logger.Log("notice", "Test")

  logger.Log("Test")
  logger.Log("Test", "Test")
  logger.Log("i", "Test")
  logger.Log("info", "Test")
  logger.Log("information", "Test")

  logger.Log("d", "Test")
  logger.Log("debug", "Test")

  // Test without colors
  logger.Colorize = false
  logger.Log("e", "Test")
  logger.Log("c", "Test")
  logger.Log("err", "Test")
  logger.Log("w", "Test")
  logger.Log("n", "Test")
  logger.Log("Test")
  logger.Log("i", "Test")
}
