package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type router struct {
	repo *repo
}

type repo struct {
	tasks []task
}

type task struct {
	id   string
	text string
}

func New() *router {
	return &router{
		repo: &repo{},
	}
}

func (r *router) unknown(w http.ResponseWriter, req *http.Request) {
	fmt.Println("unknown")
	w.WriteHeader(404)

}

func (r *router) getAll(w http.ResponseWriter, req *http.Request) {
	resp, err := json.Marshal(r.repo.tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func (r *router) createTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("created")
	//...
}

func (r *router) deleteTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("delete")
	//...
}

func main() {
	r := New()
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.unknown)
	mux.HandleFunc("/tasks", r.getAll)
	mux.HandleFunc("/task", r.createTask)
	mux.HandleFunc("/delete", r.deleteTask)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
