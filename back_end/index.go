package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT_NO = ":8080"

func main() {
	fmt.Printf("Starting application on port%s\n", PORT_NO)
	srv := &http.Server{
		Addr: PORT_NO,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
