package main

import (
  "log"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "./core"
  "./engine"
)

func main() {
  file, err := ioutil.ReadFile("./config.json")
  if err != nil {
    log.Fatal(fmt.Sprintf("%v\n", err))
  }

  c := core.Config{}
  json.Unmarshal(file, &c)

  engines := make([]*core.Engine, 0)

  for _, check := range c.Checks {
    engines = append(engines, engine.EngineFromConfig(check))
  }

  for _, engine := range engines {
    engine.Run()
  }

  select{}
}
