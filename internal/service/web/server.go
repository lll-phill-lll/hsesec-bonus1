package web

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Server interface {
	SetHandlers()
	StartServe(int) error
}

func New() Server {
	return &ServerImpl{}
}

type ServerImpl struct {
	router *mux.Router
}

func (serv *ServerImpl) SetHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/by-id", nil).Methods("GET")
	r.HandleFunc("/by-login", nil).Methods("GET")
	http.Handle("/", r)
	serv.router = r

}

func (serv *ServerImpl) StartServe(port int) error {
	portStr := strconv.Itoa(port)
	portStr = ":" + portStr

	return http.ListenAndServe(portStr, serv.router)
}
