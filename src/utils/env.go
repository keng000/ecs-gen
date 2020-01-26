package utils

// Env contains all available environment variables
// Gonna pass to envconfig
type Env struct {
	LogLevel string `envconfig:"ECS_GEN_LOG_LEVEL" default:"Info"`
}
