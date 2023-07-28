package main

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Port       string
	OutputChan StateInpChan
}

func initServer(ch StateInpChan) {
	srv := Server{
		Port:       ":3000",
		OutputChan: ch,
	}
	go srv.startServer()
}

func (srv *Server) startServer() {
	http.HandleFunc("/", midlware(srv.GetAll))
	http.HandleFunc("/insert", midlware(srv.Insert))
	http.ListenAndServe(srv.Port, nil)
}

func (srv *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	ch := make(AnswerChan)
	msg := Mess{
		ch:      ch,
		message: GetAll,
	}
	srv.OutputChan <- msg
	answ, _ := <-ch
	log.Debug("Получено сообщение из состояния: ", answ.Status)
	time.Sleep(15 * time.Second)
	JsonWriter(w, answ, http.StatusOK, "")
	log.Info("Обработан запрос")
}

func (srv *Server) Update(w http.ResponseWriter, r *http.Request) {

}

func (srv *Server) Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonWriter(w, Answer{}, http.StatusMethodNotAllowed, "MethodNotAllowed")
	}
	ch := make(AnswerChan)
	var srvData ServerData
	err := json.NewDecoder(r.Body).Decode(&srvData)
	if err != nil {
		JsonWriter(w, Answer{}, http.StatusBadRequest, err.Error())
	}
	msg := Mess{
		ch:      ch,
		message: InputData,
		data:    srvData,
	}
	log.Debug(msg.data)
	srv.OutputChan <- msg
	log.Debug("Сообщение отправлено")
	answ, _ := <-ch
	JsonWriter(w, answ, http.StatusCreated, "")
}

func JsonWriter(w http.ResponseWriter, data Answer, status int, errOutside string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil || errOutside != "" {
		errData := ErrorResponse{
			Status: false,
			Error:  err.Error(),
		}
		log.Error("server error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errData)
	}
}

func midlware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method)
		handler(w, r)
	}
}
