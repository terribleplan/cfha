package engine

import (
  "time"
  "../core"
  "../checks"
)

func createCheck(interval uint16, engine *core.Engine, host core.TargetConfig) {
  config := core.CheckCreateConfig{
    engine,
    time.Duration(int64(interval)) * time.Second,
    host,
  }
  switch host.Type {
    case "http", "https":
      checks.NewHttpChecker(config)
  }
}
