package handlers

import (
  "../cloudflare"
  "../core"
  "log"
  "fmt"
)

func NewCloudflareHandler(config core.ReactionConfig) *core.GenericHandler {
  if config.Options["email"] == "" || config.Options["apiKey"] == "" || config.Options["domain"] == "" || config.Options["name"] == "" || config.Options["ttl"] == "" {
    log.Fatal(fmt.Sprintf("Misconfigured cloudflare handler: %#v", config))
  }

  return core.NewGenericHandler(make(chan core.Transition, 5), &cloudflareHandler{
    config,
    cloudflare.NewClient(config.Options["email"], config.Options["apiKey"]),
    make(map[string]bool),
  })
}

type cloudflareHandler struct{
  config core.ReactionConfig
  client *cloudflare.Client
  actuallyDownHosts map[string]bool
}

func (this *cloudflareHandler) Handle(transition core.Transition) {
  switch transition.To {
    case core.Down:
      log.Print(fmt.Sprintf(
        "Removed cloudflare record for `%s`: `%v`\n",
        transition.RecordValue,
        this.removeCloudflareRecord(transition.RecordValue)))

    case core.Up:
      log.Print(fmt.Sprintf(
        "Added cloudflare record for `%s`: `%v`\n",
        transition.RecordValue,
        this.addCloudflareRecord(transition.RecordValue)))

    case core.Unknown: //just leave it how it was, going up/down is idempotent anyways
  }
}

func (this *cloudflareHandler) removeCloudflareRecord(recordValue string) bool {
  records, err := this.client.RetrieveRecordsByName(this.config.Options["Domain"], this.config.Options["Name"])

  if err != nil {
    return false
  }

  exitStatus := true
  for _, record := range records {
    if record.Value != recordValue {
      continue
    }

    exitStatus = exitStatus && this.client.DestroyRecord(this.config.Options["domain"], record.Id) == nil
  }

  return exitStatus
}

func (this *cloudflareHandler) addCloudflareRecord(recordValue string) bool {
  opts := cloudflare.CreateRecord{
    "A",
    this.config.Options["name"],
    recordValue,
    "1",
    "0",
  }

  _, err := this.client.CreateRecord(this.config.Options["domain"], &opts)

  if err != nil {
    return fmt.Sprintf("%s", err) == "API Error: The record already exists."
  }
  return true
}
