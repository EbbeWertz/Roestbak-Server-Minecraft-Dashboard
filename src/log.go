package main

import (
	"net/http"
)

func handleLogs(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	if n == "" {
		n = "300"
	}
	out, _ := runCmd("sudo", "journalctl", "-u", serviceName, "-n", n, "--no-pager")
	data := map[string]any{"Logs": out, "N": n}
	_ = templates.ExecuteTemplate(w, "logs.html", data)
}
