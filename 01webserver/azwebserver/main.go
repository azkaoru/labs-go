package main

import (
  "log"
  "net/http"
)

func main() {

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

  w.Write([]byte(`
      <html>
        <head>
        <title>azweberver</title>
        <body>
         Weblocome azing Git-WebServer
        </body>
      </html>
    `))
  }

  if err:= http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServer:",err)
  }
}