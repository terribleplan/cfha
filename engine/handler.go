package engine

import (
  "../core"
  "../handlers"
)

func createHandler(handler core.ReactionConfig) *core.GenericHandler {
  switch handler.Type {
    case "cloudflare":
      return handlers.NewCloudflareHandler(handler)
    case "log":
      return handlers.NewLogHandler(handler)
  }
  return nil
}
