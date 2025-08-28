package main

import (
	"html/template"
	"net/http"
)

// Template cache
var templates = template.Must(template.ParseGlob("templates/*.html"))

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Uptime": "2h 13m",        // placeholder
		"World":  "SurvivalWorld", // placeholder
	}
	templates.ExecuteTemplate(w, "home.html", data)
}

func worldsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Worlds":  []string{"SurvivalWorld", "CreativeWorld"},
		"Backups": []string{"backup-2025-08-01.zip", "backup-2025-08-15.zip"},
	}
	templates.ExecuteTemplate(w, "worlds.html", data)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Logs": []string{
			"[12:00] Server started",
			"[12:10] Player Steve joined",
			"[12:15] Player Alex left",
		},
	}
	templates.ExecuteTemplate(w, "logs.html", data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/worlds", worldsHandler)
	http.HandleFunc("/logs", logsHandler)

	println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
