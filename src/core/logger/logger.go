package logger

import (
	"github.com/hamidteimouri/go-hexagonal-architecture/core/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})

	log.SetOutput(os.Stdout)

	if config.IsDebugMode() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

