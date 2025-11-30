package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// We are serving two templates here, so we need to define a slice that stores the path for both files
	// it is important that we put the file that contains the base template *first* in the slice
	// and then each template in the order that it is defined in the base template
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// Still using the ParseFiles function, we can just pass a variadic parameter.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Using the ExecuteTemplate function instead of Execute, where we can define which template to use, "base" in this case
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := (strconv.Atoi(r.URL.Query().Get("id")))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Created new snippet"))
}
