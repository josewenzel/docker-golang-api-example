package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("{\"status\": \"I'm alive\"}"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
