package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Program go_server is starting")
	inputStateChan := make(StateInpChan)
	initServer(inputStateChan)
	initState(inputStateChan)
	time.Sleep(250 * time.Second)
}
