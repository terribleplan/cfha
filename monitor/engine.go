package monitor

import (
)

type Engine struct {
  Input chan Result
  output []GenericHandler
}

func CreateEngine(handlers []GenericHandler) *Engine {
  input := make(chan Result)

  engine := Engine{
    input,
    handlers,
  }

  e := &engine

  go e.startProcessor()

  return e
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
    for _, relay := range this.output {
      relay.channel<- change
    }

    //And set our new status
    statuses[result.RecordValue] = result.Status
  }
}
