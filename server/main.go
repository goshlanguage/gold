package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func frontpage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"templates/base.gtpl",
		"templates/navbar.gtpl",
		"templates/frontpage.gtpl",
	)
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(
			"templates/base.gtpl",
			"templates/navbar.gtpl",
			"templates/login.gtpl",
		)

		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Parsing login for: ", r.Form["username"])

	}
}


func editHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "YOOO")
}

func main() {
    // Setup in memory session store for user tokens
    sessionStore = make(map[string]Client)
    

    // Setup static router
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

    // Setup non-dynamic routes
	http.HandleFunc("/", frontpage)
	http.HandleFunc("/goldfish", ProcessLogin)
    http.HandleFunc("/edit", Authenticate(editHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
