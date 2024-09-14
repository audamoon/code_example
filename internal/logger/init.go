package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

var Logger = logrus.New()

func Init() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
}
