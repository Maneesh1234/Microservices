package main

//TO RUN WEB SERVER OPEN INTO NEXT TERMINAL
// curl -v http://localhost:9090
//go run ProductApi.go
//cd Microservices

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// reqeusts to the path /goodbye with be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// any other request will be handled by this function
	//It register the function to a path on a thing called default serve mux
	// default serve mux is http handler
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		// read the body
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)
			// // THIS METHOD SEND ERROR MESSAGE TO MY USER= 1st way
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte"Oops")

			//Second way Writing Error message to user
			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// // WRITE RESPONSE ON CONSOLE
		// log.Printf("Data is %s :", b)

		// write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	// // It construct an http server and register default handler to it
	// //second parameter is the handler
	// http.ListenAndServe(":9090", nil)

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)

}
