package main

import (
  "log"
  "fmt"
  "os"
  "syscall"
  "io/ioutil"
  "os/signal"
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

  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  <-sigs

  log.Print(fmt.Sprintf("Stopping."))
  for _, engine := range engines {
    engine.Stop()
  }
  log.Print("Exiting.")
}
