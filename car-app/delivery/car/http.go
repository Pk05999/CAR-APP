package car

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Encrypto07/car-app/datastore"
	"github.com/Encrypto07/car-app/entities"
)

type CarHandler struct {
	datastore datastore.Car
}

func New(car datastore.Car) CarHandler {
	return CarHandler{datastore: car}
}

func (a CarHandler) Hander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.get(w, r)
	case http.MethodPost:
		a.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (a CarHandler) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write([]byte("invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.datastore.Get(i)
	if err != nil {
		_, _ = w.Write([]byte("could not retrive car"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}

func (a CarHandler) create(w http.ResponseWriter, r *http.Request) {
	var car entities.Car
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := a.datastore.Create(car)
	if err != nil {
		_, _ = w.Write([]byte("could not create car"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
}

func (a CarHandler) update(w http.ResponseWriter, r *http.Request) {

}
