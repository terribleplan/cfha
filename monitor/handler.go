package monitor

type handler interface{
  handle(transition Transition)
}

type GenericHandler struct {
  channel chan Transition
}

func runHandler(input chan Transition, handler handler) *GenericHandler {
  go func() {
    for true {
      handler.handle(<-input)
    }
  }()
  return &GenericHandler{
    input,
  }
}

func CreateHandler(handler ReactionConfig) *GenericHandler {
  switch handler.Type {
    case "cloudflare":
      return newCloudflareHandler(handler)
    case "log":
      return newLogHandler(handler)
  }
  return nil
}
