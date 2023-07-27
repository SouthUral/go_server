package main

type StorageSt map[string]interface{}

type StateInpChan chan Mess

type Mess struct {
	ch      AnswerChan
	message string
	data    ServerData
}

type ServerData struct {
	key   string
	value interface{}
}

type AnswerChan chan Answer

type Answer struct {
	data   StorageSt
	status string
}
