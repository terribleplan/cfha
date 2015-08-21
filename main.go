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

  handlers := make([]monitor.GenericHandler, 2)
  handlers[0] = monitor.NewLogHandler()
  handlers[1] = monitor.NewCloudflareHandler(c.Cloudflare)

  engine := monitor.CreateEngine(handlers)

  for _, endpoint := range c.Hosts {
    monitor.CreateCheck(c.Interval, engine, endpoint)
  }

  select{}
}
