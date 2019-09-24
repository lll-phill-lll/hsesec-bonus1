package main

import (
	"fmt"
	"github.com/lll-phill-lll/hsesec/internal/service/db/postgres"
	"github.com/lll-phill-lll/hsesec/internal/service/web"
)

func main() {
	_, _ = postgres.New()
	serv := web.New()
	serv.SetHandlers()
	err := serv.StartServe(9090)
	if err != nil {
		fmt.Println(err)
	}
}
