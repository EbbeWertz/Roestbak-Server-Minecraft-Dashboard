package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

func main() {
	// Parse all templates at startup
	templates = template.Must(template.ParseFiles(
		"html/base.html",
		"html/index.html",
		"html/logs.html",
		"html/server.html",
	))

	// Routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/activate", handleActivate)
	http.HandleFunc("/deactivate", handleDeactivate)
	http.HandleFunc("/logs", handleLogs)
	http.HandleFunc("/server", handleServer)
	http.HandleFunc("/server/action", handleServerAction)

	fmt.Printf("PaperMC dashboard listening on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
