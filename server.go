package main

import (
	"encoding/json"
	"net/http"

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
	http.ListenAndServe(srv.Port, nil)
}

func (srv *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Server is working"))
	ch := make(AnswerChan)
	msg := Mess{
		ch:      ch,
		message: GetAll,
	}
	srv.OutputChan <- msg
	answ, _ := <-ch
	log.Debug("Получено сообщение из состояния: ", answ.status)
	bytes, err := json.Marshal(answ.data)
	if err != nil {
		log.Error("Ошибка преобразования в json: ", err)
		w.Write([]byte("500 Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Info("Обработан запрос")

}

func (srv *Server) Update(w http.ResponseWriter, r *http.Request) {

}

func (srv *Server) Insert(w http.ResponseWriter, r *http.Request) {

}

func midlware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method)
		handler(w, r)
	}
}
