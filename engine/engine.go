package engine

import (
  "../core"
)

func EngineFromConfig(config core.CheckConfig) *core.Engine {
  engine := core.NewEngine()

  for _, reaction := range config.Reactions {
    engine.AddHandler(createHandler(reaction))
  }

  for _, target := range config.Targets {
    engine.AddCheck(createCheck(engine, target))
  }

  return engine
}
