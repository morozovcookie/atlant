package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusAccepted)
		if _, err := writer.Write(append([]byte{}, "echo"...)); err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
