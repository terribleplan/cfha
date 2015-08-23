package core

import (
  "log"
  "fmt"
)

type Engine struct {
  Input chan Result
  handlers []*GenericHandler
  killswitch chan bool
}

func NewEngine() *Engine {
  return &Engine{
    make(chan Result),
    make([]*GenericHandler, 0),
    make(chan bool, 1),
  }
}

func (this *Engine) AddHandler(handler *GenericHandler) {
  if handler == nil {
    return
  }

  this.handlers = append(this.handlers, handler)
  log.Print(fmt.Sprintf("%v", this.handlers))
}

func (this *Engine) Run() {
  go this.startProcessor()
}

func (this *Engine) startProcessor() {
  statuses := make(map[string]Status)
  for true {
    result := <-this.Input

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
  }
}
