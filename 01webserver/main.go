package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(`
      <html>
        <head>
        <title>Weblcome golang webserver</title>
        <body>
         helloworld!!!,this page is provide by golang web server(net/http)
        </body>
      </html>
    `))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
