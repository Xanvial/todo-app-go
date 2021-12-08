package datastore

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
)

type ArrayStore struct {
	data []model.TodoData
}

func NewArrayStore() *ArrayStore {
	newData := make([]model.TodoData, 0)

	return &ArrayStore{
		data: newData,
	}
}

func (as *ArrayStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	// get completed data
	completed := make([]model.TodoData, 0)
	for _, d := range as.data {
		if d.Status {
			completed = append(completed, d)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

func (as *ArrayStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {

	// get incomplete data
	incomplete := make([]model.TodoData, 0)
	for _, d := range as.data {
		if !d.Status {
			incomplete = append(incomplete, d)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomplete)
}

func (as *ArrayStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	log.Println("ArrayStore | title:", title)
	as.data = append(as.data, model.TodoData{
		Title: title,
	})
}

func (as *ArrayStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	status, _ := strconv.ParseBool(r.FormValue("status"))

	log.Println("ArrayStore | title:", title)
	log.Println("ArrayStore | status:", status)

	for idx, d := range as.data {
		if d.Title == title {
			as.data[idx].Status = status
		}
	}
}

func (as *ArrayStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	log.Println("ArrayStore | title:", title)

	// get deleted index
	delIdx := -1
	for idx, d := range as.data {
		if d.Title == title {
			delIdx = idx
			break
		}
	}

	if delIdx != -1 {
		as.data = append(as.data[:delIdx], as.data[delIdx+1:]...)
	}
}
