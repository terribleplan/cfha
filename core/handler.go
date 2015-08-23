package core

import (
)

type Handler interface{
  Handle(transition Transition)
}

type GenericHandler struct {
  Channel chan Transition
  killswitch chan bool
}

func (this *GenericHandler) Stop() {
  this.killswitch <- true
}

func (this *GenericHandler) run(proxy Handler) {
  for true {
    var transition Transition
    select {
      case <-this.killswitch:
        return
      case transition = <-this.Channel:
        proxy.Handle(transition)
    }

  }
}

func NewGenericHandler(input chan Transition, proxy Handler) *GenericHandler {
  handler := &GenericHandler{
    input,
    make(chan bool, 1),
  }
  go handler.run(proxy)
  return handler
}
