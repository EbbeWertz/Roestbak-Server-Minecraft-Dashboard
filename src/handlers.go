package main

import (
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	idle, _ := getIdleWorlds()
	data := map[string]any{
		"Idle":       idle,
		"IdleDir":    idleDir,
		"ActiveName": activeName(),
		"LevelName":  readLevelName(),
	}
	_ = templates.ExecuteTemplate(w, "index.html", data)
}

func handleActivate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if err := activateWorld(name); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDeactivate(w http.ResponseWriter, r *http.Request) {
	if err := deactivateWorld(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
