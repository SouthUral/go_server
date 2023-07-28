package main

type StorageSt map[string]interface{}

type StateInpChan chan Mess

type Mess struct {
	ch      AnswerChan
	message string
	data    ServerData
}

type ServerData struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type AnswerChan chan Answer

type Answer struct {
	Data   StorageSt `json:"data"`
	Status string    `json:"status"`
}

type ErrorResponse struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
}
