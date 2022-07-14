package app

import (
	"encoding/json"
	"html/template"
	"l0/internal/cache"
	"l0/internal/database"
	"log"
	"net/http"
	"strings"
)

type App struct {
	database.Database
	cache.Cache
}

func (app *App) GetHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/view", app.view)
	mux.HandleFunc("/api/orders/", app.orders)

	return mux
}
func (app *App) view(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("StatusMethodNotAllowed")
		return
	}
	tmpl, err := template.
		New("index.html").
		ParseFiles("../../web/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Can not expand template")
		return
	}
	err = tmpl.Execute(w, make(map[int]struct{}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Can not expand template")
		return
	}
}
func (app *App) orders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("StatusMethodNotAllowed")
		return
	}
	id := strings.Split(r.URL.Path, "/")[3]
	data, ok := app.Cache.Get(id)
	if !ok {
		var err error
		data, err = app.Database.ReadById(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		}
		app.Cache.Add(id, data)
	}

	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
