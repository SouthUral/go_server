package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func initState(inpChan StateInpChan) {
	storage := StorageSt{
		"map":  "new",
		"grow": 12,
		"name": "Juli",
		"time": time.Now(),
	}
	st := State{
		stateStorage: storage,
		inputChanal:  inpChan,
	}
	log.Info("state inited")
	go st.StateWorker()
}

type State struct {
	stateStorage StorageSt
	inputChanal  StateInpChan
}

func (st *State) StateWorker() {
	for mess := range st.inputChanal {
		switch mess.message {
		case GetAll:
			log.Debug("получен запрос от сервера на все данные")
			aw := Answer{
				Data:   st.stateStorage,
				Status: StatusOK,
			}
			// time.Sleep(15 * time.Second)
			mess.ch <- aw
			log.Debug("данные отправлены")
		case UpdateData:
			log.Debug("получен запрос от сервера на изменение данных")
			st.stateStorage[mess.data.Key] = mess.data.Value
			aw := Answer{
				Status: StatusOK,
			}
			mess.ch <- aw
			log.Debug("данные по запросу изменены")
		case InputData:
			log.Debug("получен запрос от сервера на добавление данных")
			st.stateStorage[mess.data.Key] = mess.data.Value
			aw := Answer{
				Status: StatusOK,
			}
			mess.ch <- aw
			log.Debug("данные добавлены")
		}
	}
}
