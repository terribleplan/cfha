package core

import (
)

type Handler interface{
  Handle(transition Transition)
  Stop() bool
}

func NewGenericHandler(input chan Transition, proxy Handler) *GenericHandler {
  handler := &GenericHandler{
    input,
    make(chan bool, 0),
    make(chan bool, 0),
    proxy,
  }
  go handler.run()
  return handler
}

type GenericHandler struct {
  Channel chan Transition
  killswitch chan bool
  killresponse chan bool
  proxy Handler
}

func (this *GenericHandler) Stop() bool {
  this.killswitch<- true
  return <-this.killresponse
}

func (this *GenericHandler) run() {
  for true {
    var transition Transition
    select {
      case <-this.killswitch:
        this.killresponse<- this.proxy.Stop()
        return
      case transition = <-this.Channel:
        this.proxy.Handle(transition)
    }

  }
}
