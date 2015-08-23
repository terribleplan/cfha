package monitor

import (
  "net/http"
  "time"
  "log"
  "fmt"
)

type checkConfig struct {
  engine *Engine
  interval time.Duration
  host TargetConfig
}

type httpChecker struct {
  client *http.Client
  config checkConfig
  endpoint string
}

func (this *httpChecker) run() {
  log.Print(fmt.Sprintf("Starting: %v\n", this.config.host))
  for true {
    this.config.engine.Input<- Result{
      this.config.host.Host,
      this.check(),
    }
    time.Sleep(this.config.interval)
  }
}

func (this *httpChecker) check() Status {
  req, err := http.NewRequest("GET", this.endpoint, nil)
  if err != nil {
    log.Print(fmt.Sprintf("Stopping: %v due to http NewRequest error\n", this.config.host))
    return Unknown
  }
  req.Header.Set("Host", this.config.host.Options["hostname"])
  return this.determineHttpCheckStatus(this.client.Do(req))
}

func (this *httpChecker) determineHttpCheckStatus(res *http.Response, err error) Status {
  if err != nil || res.StatusCode != 200 {
    return Down
  }
  return Up
}

func newHttpChecker(config checkConfig) *httpChecker {
  checker := httpChecker{
    &http.Client{},
    config,
    config.host.Type + "://" + config.host.Host + "/",
  }
  go checker.run()
  return &checker
}

func CreateCheck(interval uint16, engine *Engine, host TargetConfig) {
  config := checkConfig{
    engine,
    time.Duration(int64(interval)) * time.Second,
    host,
  }
  switch host.Type {
    case "http", "https":
      newHttpChecker(config)
  }
}
