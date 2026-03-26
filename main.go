package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	// Setup logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// Generate a new UUID
	id := uuid.New()
	log.WithFields(logrus.Fields{
		"uuid": id.String(),
	}).Info("Generated new UUID")
}
