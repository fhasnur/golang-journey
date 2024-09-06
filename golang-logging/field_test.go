package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "hasnur").Info("Hello World")

	logger.WithField("username", "fhasnur").
		WithField("name", "Fandi Hasnur").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "fhasnur",
		"name":     "Fandi Hasnur",
	}).Info("Hello World")
}
