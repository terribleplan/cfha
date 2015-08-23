package engine

import (
  "../core"
)

func EngineFromConfig(config core.CheckConfig) *core.Engine {
  engine := core.NewEngine()

  for _, reaction := range config.Reactions {
    engine.AddHandler(createHandler(reaction))
  }

  createCheck(config.Interval, engine, config.Target)

  return engine
}
