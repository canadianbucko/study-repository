package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"mill/account"
	"mill/plant"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func getMinerByClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method GET!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	class := r.URL.Query().Get("class")
	if class != "" {
		structie, err := plant.Filter(class)
		if err == nil {
			marshalledStruct, err := json.Marshal(structie)
			if err != nil {
				// write header internal + write err
				w.WriteHeader(http.StatusInternalServerError)
				balls, _ := json.Marshal(NewErrDTO(err))
				w.Write(balls)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(marshalledStruct)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
		}

	}
}

func buyEquipment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method POST!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	whichEquipment := mux.Vars(r)["which"]

	switch whichEquipment {
	case "picks":
		err := plant.BuyPicks()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "vents":
		err := plant.BuyVents()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "carts":
		err := plant.BuyCarts()
		if err != nil {
			/// dto
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func buyMiner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method POST!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	whichMiner := mux.Vars(r)["which"]

	switch whichMiner {
	case "small":
		id, err := plant.HireSmallMiner()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
			return
		}
		msg := "hey, there's your new Small miner, his ID -> " + strconv.Itoa(id)
		b, err := json.Marshal(NewResponseDTO(msg))
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	case "normal":
		id, err := plant.HireNormalMiner()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
			return
		}
		msg := "hey, there's your new Normal miner, his ID -> " + strconv.Itoa(id)
		b, err := json.Marshal(NewResponseDTO(msg))
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	case "super":
		id, err := plant.HireSuperMiner()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			balls, _ := json.Marshal(NewErrDTO(err))
			w.Write(balls)
			return
		}
		msg := "hey, there's your new Super miner, his ID -> " + strconv.Itoa(id)
		b, err := json.Marshal(NewResponseDTO(msg))
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
		return
	}
}

var mtx sync.Mutex

func getCurrentWorkers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method GET!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	mtx.Lock()
	defer mtx.Unlock() // to-do захуярить это нахуй и сделать функции в plant а тут их просто вызывать
	bytesMap, _ := json.Marshal(plant.MinerMap)

	w.WriteHeader(http.StatusOK)
	w.Write(bytesMap)

}

func getDeadWorkers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method GET!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	mtx.Lock() // to-do захуярить это нахуй и сделать функции в plant а тут их просто вызывать
	defer mtx.Unlock()
	bytesDeadMiners, _ := json.Marshal(plant.DeadMiners)

	w.WriteHeader(http.StatusOK)
	w.Write(bytesDeadMiners)

}

func currentBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method GET!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}

	balance := account.CurrentBalance()

	bytes, _ := json.Marshal(NewResponseDTO(balance))
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

func boughtEquipment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method GET!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	whichEquipment := mux.Vars(r)["which"]

	dtoStruct := plant.NewEquipDTO(whichEquipment)
	b, err := json.Marshal(dtoStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getMinerInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := errors.New("sorry, but you should use method GET!")
		b, _ := json.Marshal(NewErrDTO(err))
		w.Write(b)
		return
	}
	whichMinerById := mux.Vars(r)["id"]
	balls, _ := strconv.Atoi(whichMinerById) // если ошибка передали хуйню
	shitStruct := plant.MinerInfo(balls)     // и получаем в ответ MinerInfo
	byte, _ := json.Marshal(shitStruct)
	w.WriteHeader(http.StatusOK)
	w.Write(byte)
}

func startPlant(w http.ResponseWriter, r *http.Request) {
	go func() {
		plant.StartPlant()
	}()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ну да блять отьебись?"))
}
func killPlant(w http.ResponseWriter, r *http.Request) {
	plant.KillPlant()
	w.WriteHeader(http.StatusOK)
}

func StartServer() {
	fmt.Println("server is running!")
	router := mux.NewRouter()
	router.Path("/start").Methods("GET").HandlerFunc(startPlant)
	router.Path("/stop").HandlerFunc(killPlant)
	router.Path("/miners/{id}").Methods("GET").HandlerFunc(getMinerInfo)
	router.Path("/miners/buy/{which}").Methods("POST").HandlerFunc(buyMiner)
	router.Path("/equipment/buy/{which}").Methods("POST").HandlerFunc(buyEquipment)
	router.Path("/balance").Methods("GET").HandlerFunc(currentBalance)
	router.Path("/equipment/{which}").Methods("GET").HandlerFunc(boughtEquipment)
	router.Path("/currentminers/").Queries("minerclass", "small").Methods("GET").HandlerFunc(getMinerByClass)
	router.Path("/currentminers").Methods("GET").HandlerFunc(getCurrentWorkers)
	router.Path("/deadminers").Methods("GET").HandlerFunc(getDeadWorkers)

	if err := http.ListenAndServe(":9999", router); err != nil {
		fmt.Println("error starting http server", err)
	}

}

// what I want t say s sofjasdfidasfijasigit

/* шо имеем
++StartPlant() router.Path("/start").Methods("POST").HandlerFunc(start)
++KillPlant() router.Path("/stop").Methods("POST").HandlerFunc(stop)
++Println(DeadMiners) --отработавшие майнеры  router.Path("/deadminers").Methods("GET").HandlerFunc(deadminers)
++Println(MinerMap) -- работающие майнеры router.Path("/currentminers").Methods("GET").HandlerFunc(currentminers)
++Filter("string") структуру + ошибку - работающие + фильтр по классу router.Path("/currentminers/?minerclass=broke").Methods("GET").HandlerFunc(equipment)
++BuyPicks() - err 	router.Path("/equipment/buy/{picks/vents/carts}").Methods("POST?").HandlerFunc(equipment)
++BuyVents() - err 	router.Path("/equipment/buy/{picks/vents/carts}").Methods("POST?").HandlerFunc(equipment)
++BuyCarts() - err 	router.Path("/equipment/buy/{picks/vents/carts}").Methods("POST?").HandlerFunc(equipment)
++NewEquipDTO BoughtEquipment() - возвращает структуру c купленным оборудованием + ценой router.Path("/equipment/info/{"which"}").Methods("GET").HandlerFunc(equipmentiNFO)
++account.CurrentBalance() int -  возвращает баланс router.Path("/equipment/buy/{picks/vents/carts}").Methods("POST?").HandlerFunc(equipment)
++MinerInfo(id) - стурктуру с майнером  router.Path("/miners/{id}").Methods("POST?").HandlerFunc(equipment)
++BuyMiner("small") //	router.Path("/miner/{small}").Methods("POST?").HandlerFunc(miner)
*/

// посчитать количество строк в проекте всего
