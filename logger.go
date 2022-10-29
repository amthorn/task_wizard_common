package task_wizard_common

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	if os.Getenv("LOG_LEVEL") != "" {
		level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
		if err != nil { panic(err) }
		log.SetLevel(level)
		log.Debugf("Initialized logger at level '%s'", level)
	}else{
		log.SetLevel(log.InfoLevel)
		log.Debug("Initialized logger at default level INFO")
	}
}
