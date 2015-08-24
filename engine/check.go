package engine

import (
  "log"
  "fmt"
  "time"
  "../core"
  "../checks"
)

func createCheck(engine *core.Engine, host core.TargetConfig) *core.GenericCheck {
  log.Print("createCheck called")
  log.Print(fmt.Sprintf("creating a `%s`", host.Type))
  config := core.CheckCreateConfig{
    engine,
    time.Duration(int64(host.Interval)) * time.Second,
    host,
  }
  switch host.Type {
    case "http", "https":
      return checks.NewHttpChecker(config)
  }
  return nil
}
