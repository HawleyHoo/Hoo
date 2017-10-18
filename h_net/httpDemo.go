package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main()  {
	http.Handle("/", Hey)
	http.ListenAndServe(":9091", nil)

}

func Hey(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl1)
		t.Execute(w, nil)
	} else {
		//r.ParseForm()
		fmt.Println(r.FormValue("username"))
	}
}

const tpl1 = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hey</title>
</head>
<body>
<form method="post" action="/">
    UserName: <input type="text" name="username">
    Password: <input type="password", name="password">
    <button type="submit">Submit</button>
</form>
</body>
</html>
`