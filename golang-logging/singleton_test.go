package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello World")
	logrus.Warn("Hello World")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello World")
	logrus.Warn("Hello World")
}
