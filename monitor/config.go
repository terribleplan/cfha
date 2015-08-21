package monitor

type HostConfig struct {
  Host string
  Type string
  Options map[string]string
}

type CloudflareConfig struct {
  Email string
  ApiKey string
  Domain string
  Name string
  Ttl string
}

type Config struct {
  Cloudflare CloudflareConfig
  Hosts []HostConfig
  Interval uint16
}
