package monitor

import (
  "fmt"
  "log"
)

func newLogHandler(config ReactionConfig) *GenericHandler {
  return runHandler(make(chan Transition), &logHandler{})
}

type logHandler struct{}

func (this *logHandler) handle(transition Transition) {
  log.Print(fmt.Sprintf(
    "`%s` has become `%d` - `%s`",
    transition.RecordValue, transition.To,
    transition.To.String()))
}
