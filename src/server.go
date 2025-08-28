package main

import (
	"net/http"
)

func handleServer(w http.ResponseWriter, r *http.Request) {
	status, _ := runCmd("systemctl", "is-active", serviceName)
	data := map[string]any{"Status": status}
	_ = templates.ExecuteTemplate(w, "server.html", data)
}

func handleServerAction(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("action")
	switch a {
	case "start", "stop", "restart":
		_, _ = runCmd("sudo", "systemctl", a, serviceName)
	}
	http.Redirect(w, r, "/server", http.StatusSeeOther)
}
