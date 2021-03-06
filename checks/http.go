package checks

import (
  "net/http"
  "../core"
  "time"
  "log"
  "fmt"
)

func NewHttpChecker(config core.CheckCreateConfig) *core.GenericCheck {
  return core.NewGenericCheck(&HttpChecker{
    &http.Client{},
    config,
    config.Host.Type + "://" + config.Host.Host + "/",
  })
}

type HttpChecker struct {
  client *http.Client
  config core.CheckCreateConfig
  endpoint string
}

func (this *HttpChecker) Check() time.Duration {
  this.config.Engine.Input<- core.Result{
    this.config.Host.Host,
    this.check(),
  }
  return this.config.Interval
}

func (this *HttpChecker) Stop() bool {
  return true
}

func (this *HttpChecker) check() core.Status {
  req, err := http.NewRequest("GET", this.endpoint, nil)
  if err != nil {
    log.Print(fmt.Sprintf("Stopping: %v due to http NewRequest error\n", this.config))
    return core.Unknown
  }
  req.Header.Set("Host", this.config.Host.Options["hostname"])
  return this.determineHttpCheckStatus(this.client.Do(req))
}

func (this *HttpChecker) determineHttpCheckStatus(res *http.Response, err error) core.Status {
  if err != nil || res.StatusCode != 200 {
    return core.Down
  }
  return core.Up
}
