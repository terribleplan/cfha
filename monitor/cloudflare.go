package monitor

import (
  "../cloudflare"
  "log"
  "fmt"
)

func NewCloudflareHandler(config CloudflareConfig) GenericHandler {
  return runHandler(make(chan Transition, 5), &cloudflareHandler{
    config,
    cloudflare.NewClient(config.Email, config.ApiKey),
    make(map[string]bool),
  })
}

type cloudflareHandler struct{
  config CloudflareConfig
  client *cloudflare.Client
  actuallyDownHosts map[string]bool
}

func (this *cloudflareHandler) handle(transition Transition) {
  switch transition.To {
    case Down:
      log.Print(fmt.Sprintf(
        "Removed cloudflare record for `%s`: `%v`\n",
        transition.RecordValue,
        removeCloudflareRecord(this.client, this.config, transition.RecordValue)))

    case Up:
      log.Print(fmt.Sprintf(
        "Added cloudflare record for `%s`: `%v`\n",
        transition.RecordValue,
        addCloudflareRecord(this.client, this.config, transition.RecordValue)))

    case Unknown: //just leave it how it was, going up/down is idempotent anyways
  }
}

func removeCloudflareRecord(client *cloudflare.Client, config CloudflareConfig, recordValue string) bool {
  records, err := client.RetrieveRecordsByName(config.Domain, config.Name)

  if err != nil {
    return false
  }

  exitStatus := true
  for _, record := range records {
    if record.Value != recordValue {
      continue
    }

    exitStatus = exitStatus && client.DestroyRecord(config.Domain, record.Id) == nil
  }

  return exitStatus
}

func addCloudflareRecord(client *cloudflare.Client, config CloudflareConfig, recordValue string) bool {
  opts := cloudflare.CreateRecord{
    "A",
    config.Name,
    recordValue,
    "1",
    "0",
  }

  _, err := client.CreateRecord(config.Domain, &opts)

  if err != nil {
    return fmt.Sprintf("%s", err) == "API Error: The record already exists."
  }
  return true
}
