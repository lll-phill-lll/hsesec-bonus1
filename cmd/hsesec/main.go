package main

import (
	"fmt"
	"github.com/lll-phill-lll/hsesec/internal/service/db"
	"github.com/lll-phill-lll/hsesec/internal/service/web"
)

func main() {
	database, err := db.New()
	if err != nil {
		fmt.Println(err)
	}
	serv := web.New(database)
	serv.SetHandlers()
	err = serv.StartServe(9090)
	if err != nil {
		fmt.Println(err)
	}
}
