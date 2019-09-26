package web

import (
	"fmt"
	"github.com/lll-phill-lll/hsesec/internal/service/db/postgres"
	"net/http"
	"strconv"
)

const usersResponseTemplate = "<html> <head> <title>Users</title> </head> <body> %s </body> </html>"

func (serv *ServerImpl) allUsers(w http.ResponseWriter, r *http.Request) {
	users, err := serv.DB.LoadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := constructUsersTable(users)
	fmt.Println(resp)
	_, err = fmt.Fprintln(w, resp)
	if err != nil {
		fmt.Println(err)
	}
}

func (serv *ServerImpl) byLogin(w http.ResponseWriter, r *http.Request) {
	logins, ok := r.URL.Query()["login"]
	if !ok || len(logins) < 1 {
		fmt.Println("not found")
		return
	}
	users, err := serv.DB.LoadByLogin(logins[0])
	if err != nil {
		fmt.Println("not found")
		return
	} else {
		_, err := fmt.Fprintln(w, constructUsersTable(users))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (serv *ServerImpl) byID(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		fmt.Println("not found")
		return
	}
	id, err := strconv.Atoi(ids[0])
	users, err := serv.DB.LoadByID(id)
	if err != nil {
		fmt.Println("not found")
		return
	} else {
		_, err := fmt.Fprintln(w, constructUsersTable(users))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func constructUsersTable(users []postgres.User) string {
	str := `<table border="1">`
	str += "<tr>"
	str += fmt.Sprintf("<td>#</td>")
	str += fmt.Sprintf("<td>ID</td>")
	str += fmt.Sprintf("<td>Login</td>")
	str += fmt.Sprintf("<td>Card number</td>")
	str += fmt.Sprintf("<td>Money amount</td>")
	str += fmt.Sprintf("<td>Status (active)</td>")
	str += "</tr>"
	for i, user := range users {
		str += "<tr>"
		str += fmt.Sprintf("<td>%d</td>", i+1)
		str += fmt.Sprintf("<td>%d</td>", user.Id)
		str += fmt.Sprintf("<td>%s</td>", user.Login)
		str += fmt.Sprintf("<td>%s</td>", user.CardNumber)
		str += fmt.Sprintf("<td>%d</td>", user.MoneyAmount)
		str += fmt.Sprintf("<td>%t</td>", user.Status)
		str += "</tr>"

	}
	str += "</table>"
	return fmt.Sprintf(usersResponseTemplate, str)
}
