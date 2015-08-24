package core

import (
  "log"
  "fmt"
)

type Engine struct {
  Input chan Result
  handlers []*GenericHandler
  checks []*GenericCheck
  killswitch chan bool
  killresponse chan bool
}

func NewEngine() *Engine {
  engine := &Engine{
    make(chan Result),
    make([]*GenericHandler, 0),
    make([]*GenericCheck, 0),
    make(chan bool, 0),
    make(chan bool, 0),
  }

  go engine.startProcessor()

  return engine
}

func (this *Engine) AddCheck(check *GenericCheck) {
  if check == nil {
    return
  }

  this.checks = append(this.checks, check)
  log.Print(fmt.Sprintf("%v", this.checks))
}

func (this *Engine) AddHandler(handler *GenericHandler) {
  if handler == nil {
    return
  }

  this.handlers = append(this.handlers, handler)
  log.Print(fmt.Sprintf("%v", this.handlers))
}

func (this *Engine) Stop() bool {
  this.killswitch<- true
  return <-this.killresponse
}

func (this *Engine) startProcessor() {
  statuses := make(map[string]Status)
  for true {
    select {
      case result := <-this.Input:
        //No transition if we don't exist
        if result.Status == statuses[result.RecordValue] {
          continue
        }

        //Create a record with to, from
        change := Transition{
          result.Status,
          statuses[result.RecordValue],
          result.RecordValue,
        }

        //Send the record to everyone who cares
        for _, relay := range this.handlers {
          relay.Channel<- change
        }

        //And set our new status
        statuses[result.RecordValue] = result.Status
      case <-this.killswitch:
        this.killresponse<- this.stop()
        return
    }
  }
}

func (this *Engine) stop() bool {
  exitStatus := true
  for _, handler := range this.handlers {
    exitStatus = exitStatus && handler.Stop()
  }
  for _, check := range this.checks {
    exitStatus = exitStatus && check.Stop()
  }
  return exitStatus
}

