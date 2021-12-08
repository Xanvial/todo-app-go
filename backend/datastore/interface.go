package datastore

import "net/http"

type DataStore interface {
	GetCompleted(w http.ResponseWriter, r *http.Request)
	GetIncomplete(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
