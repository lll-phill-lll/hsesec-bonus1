package web

import (
	"fmt"
	"net/http"
)

const usersResponeTemplate = "<html> <head> <title>Users</title> </head> <body> %s </body> </html>"

func (serv *ServerImpl) allUsers(w http.ResponseWriter, r *http.Request) {
	str := `<table border="1">`
	users, err := serv.DB.LoadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
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
	resp := fmt.Sprintf(usersResponeTemplate, str)
	fmt.Println(resp)
	_, err = fmt.Fprintln(w, resp)
	if err != nil {
		fmt.Println(err)
	}
}
