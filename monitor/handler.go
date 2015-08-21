package monitor

type handler interface{
  handle(transition Transition)
}

type GenericHandler struct {
  channel chan Transition
}

func runHandler(input chan Transition, handler handler) GenericHandler {
  go func() {
    for true {
      handler.handle(<-input)
    }
  }()
  return GenericHandler{
    input,
  }
}
