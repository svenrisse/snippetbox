package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.logRequest(secureHeaders(mux))
}
