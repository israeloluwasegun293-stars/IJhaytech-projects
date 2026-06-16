package handlers

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost{
		w.Header().Set("Allow method Get", http.MethodGet)

		http.Error(w, "Method not alowed", 405)
		return
	}

	temp, err := template.ParseFiles("templates/home.html")
	if err != nil{
		http.Error(w, "internal server error", 500)
	}
	temp.Execute(w, nil)
}