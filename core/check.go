package core

import (
  "time"
)

type CheckCreateConfig struct {
  Engine *Engine
  Interval time.Duration
  Host TargetConfig
}
