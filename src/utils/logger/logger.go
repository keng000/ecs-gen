package logger

import (
	"fmt"
	"os"

	"github.com/keng000/ecs-gen/src/utils"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// SetupLogger is an initial process of logrus
func SetupLogger() {
	logrusConfig()
}

func logrusConfig() {
	env := &utils.Env{}
	if err := envconfig.Process("", env); err != nil {
		logrus.WithFields(logrus.Fields{
			"env": env,
		}).Panic("[PANIC] failed to get env var")

	}

	switch env.LogLevel {
	case "Error", "ERROR", "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "Warn", "WARN", "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "Info", "INFO", "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "Debug", "DEBUG", "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "Trace", "TRACE", "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(os.Stdout)

	logrus.WithFields(logrus.Fields{
		"env": fmt.Sprintf("%+v", env),
	}).Debug("List envs")
}
