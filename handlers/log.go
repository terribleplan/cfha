package handlers

import (
  "fmt"
  "log"
  "../core"
)

func NewLogHandler(config core.ReactionConfig) *core.GenericHandler {
  return core.NewGenericHandler(make(chan core.Transition), &logHandler{})
}

type logHandler struct{}

func (this *logHandler) Handle(transition core.Transition) {
  log.Print(fmt.Sprintf(
    "`%s` has become `%d` - `%s`",
    transition.RecordValue, transition.To,
    transition.To.String()))
}
