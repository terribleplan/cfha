package core

type Config struct {
  Checks []CheckConfig
}

type CheckConfig struct {
  Targets []TargetConfig
  Reactions []ReactionConfig
}

type TargetConfig struct {
  Interval uint16
  Type string
  Host string
  Options map[string]string
}

type ReactionConfig struct {
  Type string
  Options map[string]string
}
