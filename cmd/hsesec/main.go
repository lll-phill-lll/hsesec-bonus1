package main

import (
	"fmt"
	db2 "github.com/lll-phill-lll/hsesec/internal/service/db"
	"github.com/lll-phill-lll/hsesec/internal/service/web"
)

func main() {
	db, err := db2.New()
	if err != nil {
		fmt.Println(err)
	}
	serv := web.New(db)
	serv.SetHandlers()
	err = serv.StartServe(9090)
	if err != nil {
		fmt.Println(err)
	}
}
