package core

type Config struct {
  Checks []CheckConfig
}

type CheckConfig struct {
  Interval uint16
  Target TargetConfig
  Reactions []ReactionConfig
}

type TargetConfig struct {
  Type string
  Host string
  Options map[string]string
}

type ReactionConfig struct {
  Type string
  Options map[string]string
}
