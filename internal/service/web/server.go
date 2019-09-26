package web

import (
	"github.com/gorilla/mux"
	"github.com/lll-phill-lll/hsesec/internal/service/db"
	"net/http"
	"strconv"
)

type Server interface {
	SetHandlers()
	StartServe(int) error
}

func New(base db.DataBase) Server {
	return &ServerImpl{DB: base}
}

type ServerImpl struct {
	router *mux.Router
	DB     db.DataBase
}

func (serv *ServerImpl) SetHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/users", serv.allUsers).Methods("GET")
	r.HandleFunc("/by-id/", serv.byID).Methods("GET")
	r.HandleFunc("/by-login/", serv.byLogin).Methods("GET")
	r.HandleFunc("/", serv.home).Methods("GET")
	http.Handle("/", r)
	serv.router = r

}

func (serv *ServerImpl) StartServe(port int) error {
	portStr := strconv.Itoa(port)
	portStr = ":" + portStr

	return http.ListenAndServe(portStr, serv.router)
}
