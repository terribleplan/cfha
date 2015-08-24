package core

import (
  "time"
)

type CheckCreateConfig struct {
  Engine *Engine
  Interval time.Duration
  Host TargetConfig
}

type Check interface{
  Check() time.Duration
  Stop() bool
}

func NewGenericCheck(proxy Check) *GenericCheck {
  check := &GenericCheck{
    make(chan bool, 0),
    make(chan bool, 0),
    proxy,
  }
  go check.run()
  return check
}

type GenericCheck struct {
  killswitch chan bool
  killresponse chan bool
  proxy Check
}

func (this *GenericCheck) Stop() bool {
  this.killswitch<- true
  return <-this.killresponse
}

func (this *GenericCheck) run() {
  timeout := time.NewTimer(0)
  for true {
    select {
      case <-this.killswitch:
        this.killresponse<- this.proxy.Stop()
        return
      case <-timeout.C:
        timeout = time.NewTimer(this.proxy.Check())
    }
  }
}
