package web

import (
	"fmt"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	a := `<html>
 <head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>Тег TABLE</title>
 </head>
 <body>
  <table border="1" width="100%" cellpadding="5">
   <tr>
    <th>Ячейка 1</th>
    <th>Ячейка 2</th>
   </tr>
   <tr>
    <td>Ячейка 3</td>
    <td>Ячейка 4</td>
  </tr>
 </table>
 </body>`
	_, err := fmt.Fprintln(w, a)
	if err != nil {
		fmt.Println(err)
	}
}
