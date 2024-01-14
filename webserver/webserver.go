package webserver

import "net/http"

func MyWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3002", nil)
}

// el * indica que es un puntero
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
