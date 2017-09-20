package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type usersMap struct {
	sync.RWMutex
	users map[uint]User // users by ID
}

var users usersMap

func main() {
	users = usersMap{
		users: make(map[uint]User),
	}

	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/set", setHandler).Methods("POST")
	r.HandleFunc("/delete", deleteHandler).Methods("DELETE")
	fmt.Println(http.ListenAndServe("0.0.0.0:8088", r))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idnum, err := strconv.Atoi(id)
	if err != nil || idnum < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users.RLock()
	defer users.RUnlock()
	user, ok := users.users[uint(idnum)]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	d, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(d)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users.Lock()
	defer users.Unlock()
	users.users[user.ID] = user
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idnum, err := strconv.Atoi(id)
	if err != nil || idnum < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users.Lock()
	defer users.Unlock()
	if _, ok := users.users[uint(idnum)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(users.users, uint(idnum))
}
