package wbserver

import "net/http"

func MyWebServer() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/wbserver/index.html")
}
