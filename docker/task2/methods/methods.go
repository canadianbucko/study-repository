package methods

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"yayca/dto"
	"yayca/middleware"
	"yayca/sql"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
)

// ну тут нужно описать методы /
//  1. `POST /employees` — Создать нового работника в кадровом учёте
//  2. `GET /employees` — Получить список всех работников из кадрового учёта
//  3. `DELETE /employees` — Удалить какого-то работника из кадрового учёта
//
// то есть я получаю там что ну да получать буду сам джсон или еще че-нибудь в нем и просто парсить в дто правильно понимаю?
// postman - json (full name, position) - unmarshall - to dto - ну да че-то такое
var employeeMap = make(map[int]dto.EmployeeDTO) // тут создаю мапу, теперь в нее нужно напихать значения!
var mtx sync.Mutex

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("sorry, this method is not allowed")
		middleware.SendError(w, http.StatusMethodNotAllowed, err)
		return
	}

	var employee dto.EmployeeDTO // это сюда мы положили наших работчников

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		err := errors.New("something wrong with json")
		middleware.SendError(w, http.StatusBadRequest, err)
		return
	}
	newEmployee, err := dto.NewEmployeeDTO(employee.FullName, employee.Position) // тут вызываю конструктор
	if err != nil {
		middleware.SendError(w, http.StatusBadRequest, err)
		return
	}

	// ?? connect to the database
	if err := sql.ShoveThatIntoSQL(newEmployee); err != nil {
		pp.Println(err)
		middleware.SendError(w, http.StatusInternalServerError, err)
		return
	}

	// если мы дошли до сюда значит всё ок и вся валидация пройдена
	w.WriteHeader(http.StatusOK)

	mtx.Lock()
	defer mtx.Unlock()
	employeeMap[newEmployee.ID] = newEmployee // и вот мы положили эту красавицу вот сюда! ура вот наш работничек!

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New("well, you should you another method (GET) in this case.")
		middleware.SendError(w, http.StatusMethodNotAllowed, err)
		return
	}
	//  смылса нет делать временную мапу и перегонять. только если хотим скрыть поля (password)

	mtx.Lock()
	defer mtx.Unlock() //rlock /runlock лучше тут ток чтение!
	responseMapFinal, _ := json.Marshal(employeeMap)
	w.WriteHeader(http.StatusOK)
	w.Write(responseMapFinal)
	// ?
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		err := errors.New("sorry, use method (DELETE)")
		middleware.SendError(w, http.StatusMethodNotAllowed, err)
		return
	}
	id := r.URL.Query().Get("id") // /employees?id=2
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("bad request, what the fuck have you send?")
		middleware.SendError(w, http.StatusBadRequest, err)
		return
	}
	_, ok := employeeMap[idInt] //проверяем есть ли в мапе такое value, ok := employeeMap[idInt]
	if !ok {
		err := errors.New("nein, nicht this id, fuck off!")
		middleware.SendError(w, http.StatusNotFound, err)
		return
	}
	delete(employeeMap, idInt)
	w.WriteHeader(http.StatusOK)

}

func StartServer() {
	pp.Println("preparing to start the server...")
	router := mux.NewRouter()

	router.Path("/employees").Methods("POST").HandlerFunc(PostEmployee)
	router.Path("/employees").Methods("GET").HandlerFunc(GetEmployees)
	router.Path("/employees").Methods("DELETE").HandlerFunc(DeleteEmployee)

	if err := http.ListenAndServe(":9999", router); err != nil {
		panic(err)
	}

}

/*
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		msg := "error decoding a json message " + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}
		this is a sample text. please ignore it
*/
