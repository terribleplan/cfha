package main

import (
  "log"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "./monitor"
)

func main() {
  file, err := ioutil.ReadFile("./config.json")
  if err != nil {
    log.Fatal(fmt.Sprintf("%v\n", err))
  }

  c := monitor.Config{}
  json.Unmarshal(file, &c)

  for _, check := range c.Checks {
    handlers := make([]*monitor.GenericHandler, 0)

    for _, reaction := range check.Reactions {
      handler := monitor.CreateHandler(reaction)

      if handler == nil {
        continue
      }

      handlers = append(handlers, handler)
    }

    engine := monitor.CreateEngine(handlers)
    monitor.CreateCheck(check.Interval, engine, check.Target)
  }

  select{}
}
