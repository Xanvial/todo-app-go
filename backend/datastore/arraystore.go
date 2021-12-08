package datastore

import (
	"encoding/json"
	"net/http"

	"github.com/Xanvial/todo-app-go/model"
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
