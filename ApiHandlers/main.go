package main

//go run main.go
//cd Microservices
//CD handler
// import (
// 	"Microservices/handlers"
// 	"time"

// 	"log"
// 	"net/http"
// 	"os"
// )

// // var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

// func main() {

// 	// env.Parse()

// 	// LOGGER IS FOR OUTPUT
// 	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

// 	// create the handlers
// 	hh := handlers.NewHello(l)
// 	gh := handlers.NewGoodbye(l)

// 	//REGISTER MY HANDLER WITH SERVER
// 	//ServeMux is the type of object
// 	//SERVER HAVE DEFAULT HANDLER AND THIS HANDLER IS THE HTTP ServeMux
// 	// create a new serve mux and register the handlers
// 	//CREATE A NEW serve mux
// 	sm := http.NewServeMux()
// 	//register the handlers
// 	sm.Handle("/", hh)
// 	sm.Handle("/goodbye", gh)

// 	// FIRST PARAM IS BINDING ADDRESS AND SECOND IS HTTP HANDLER
// 	// WHEN SECOND PARAM IS NIL THEN SERVER PROVIDE DEFAULT SERVE MUX
// 	//SERVE MUX ALSO IMPLEMENT HANDLER INTERFACE
// 	// 1st way
// 	// http.ListenAndServe("9090:", sm)
// 	//2nd way
// 	log.Println("Starting Server")
// 	err := http.ListenAndServe(":9090", sm)
// 	log.Fatal(err)

// 	//SOMETIME LISTEN AND SERVE DENIAL THE REQUEST WHEN MANY REQUESST COMES
// 	//SO OVERCOMES THIS ISSUES BY CREATING A SERVER
// 	// create a new server
// 	// s := &http.Server{
// 	// 	// Addr:         *bindAddress,      // configure the bind address
// 	// 	Addr:         9090,
// 	// 	Handler:      sm,                // set the default handler
// 	// 	ErrorLog:     l,                 // set the logger for the server
// 	// 	ReadTimeout:  5 * time.Second,   // max time to read request from the client
// 	// 	WriteTimeout: 10 * time.Second,  // max time to write response to the client
// 	// 	IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
// 	// }

// 	// s.ListenAndServe()

// 	// // start the server
// 	// go func() {
// 	// 	l.Println("Starting server on port 9090")

// 	// 	err := s.ListenAndServe()
// 	// 	if err != nil {
// 	// 		l.Printf("Error starting server: %s\n", err)
// 	// 		os.Exit(1)
// 	// 	}
// 	// }()

// 	// // trap sigterm or interupt and gracefully shutdown the server
// 	// c := make(chan os.Signal, 1)
// 	// signal.Notify(c, os.Interrupt)
// 	// signal.Notify(c, os.Kill)

// 	// // Block until a signal is received.
// 	// sig := <-c
// 	// log.Println("Got signal:", sig)

// 	// // gracefully shutdown the server, waiting max 30 seconds for current operations to complete
// 	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	// s.Shutdown(ctx)
// }

// // (2)
// import (
// 	"Microservices/handlers"
// 	"time"

// 	"log"
// 	"net/http"
// 	"os"
// )

// // var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

// func main() {

// 	// env.Parse()

// 	// LOGGER IS FOR OUTPUT
// 	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

// 	// create the handlers
// 	hh := handlers.NewHello(l)
// 	gh := handlers.NewGoodbye(l)

// 	//REGISTER MY HANDLER WITH SERVER
// 	// Handler is the interface
// 	//ServeMux is the type of object and it is http handler
// 	//SERVER HAVE DEFAULT HANDLER AND THIS HANDLER IS THE HTTP ServeMux
// 	// create a new serve mux and register the handlers
// 	//CREATE A NEW serve mux
// 	sm := http.NewServeMux()
// 	//register the handlers
// 	sm.Handle("/", hh)
// 	sm.Handle("/goodbye", gh)

// 	// FIRST PARAM IS BINDING ADDRESS AND SECOND IS HTTP HANDLER
// 	// WHEN SECOND PARAM IS NIL THEN SERVER PROVIDE DEFAULT SERVE MUX
// 	//SERVE MUX ALSO IMPLEMENT HANDLER INTERFACE
// 	// 1st way
// 	// http.ListenAndServe("9090:", sm)
// 	//2nd way
// 	// log.Println("Starting Server")
// 	// err := http.ListenAndServe(":9090", sm)
// 	// log.Fatal(err)

// 	//SOMETIME LISTEN AND SERVE DENIAL THE REQUEST WHEN MANY REQUESST COMES
// 	//SO OVERCOMES THIS ISSUES BY CREATING A SERVER
// 	// create a new server
// 	s := &http.Server{
// 		// Addr:         *bindAddress,      // configure the bind address
// 		Addr:         ":9090",
// 		Handler:      sm,                // set the default handler
// 		ErrorLog:     l,                 // set the logger for the server
// 		ReadTimeout:  5 * time.Second,   // max time to read request from the client
// 		WriteTimeout: 10 * time.Second,  // max time to write response to the client
// 		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
// 	}

// 	s.ListenAndServe()

// 	// // start the server
// 	// go func() {
// 	// 	l.Println("Starting server on port 9090")

// 	// 	err := s.ListenAndServe()
// 	// 	if err != nil {
// 	// 		l.Printf("Error starting server: %s\n", err)
// 	// 		os.Exit(1)
// 	// 	}
// 	// }()

// 	// // trap sigterm or interupt and gracefully shutdown the server
// 	// c := make(chan os.Signal, 1)
// 	// signal.Notify(c, os.Interrupt)
// 	// signal.Notify(c, os.Kill)

// 	// // Block until a signal is received.
// 	// sig := <-c
// 	// log.Println("Got signal:", sig)

// 	// // gracefully shutdown the server, waiting max 30 seconds for current operations to complete
// 	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	// s.Shutdown(ctx)
// }

//(3)
// (2)
import (
	"Microservices/handlers"
	"os/signal"

	// "std/context"
	"context"
	"time"

	"log"
	"net/http"
	"os"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	// env.Parse()

	// LOGGER IS FOR OUTPUT
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	//REGISTER MY HANDLER WITH SERVER
	// Handler is the interface
	//ServeMux is the type of object and it is http handler
	//SERVER HAVE DEFAULT HANDLER AND THIS HANDLER IS THE HTTP ServeMux
	// create a new serve mux and register the handlers
	//CREATE A NEW serve mux
	sm := http.NewServeMux()
	//register the handlers
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// FIRST PARAM IS BINDING ADDRESS AND SECOND IS HTTP HANDLER
	// WHEN SECOND PARAM IS NIL THEN SERVER PROVIDE DEFAULT SERVE MUX
	//SERVE MUX ALSO IMPLEMENT HANDLER INTERFACE
	// 1st way
	// http.ListenAndServe("9090:", sm)
	//2nd way
	// log.Println("Starting Server")
	// err := http.ListenAndServe(":9090", sm)
	// log.Fatal(err)

	//SOMETIME LISTEN AND SERVE DENIAL THE REQUEST WHEN MANY REQUESST COMES
	//SO OVERCOMES THIS ISSUES BY CREATING A SERVER
	// create a new server
	s := &http.Server{
		// Addr:         *bindAddress,      // configure the bind address
		Addr:         ":9090",
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// s.ListenAndServe()

	// // start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	// IT WILL BROADCOST A MESSAGE TO THE CHANNEL WHEN OS KILL OR INTERRUPT COMMAND TAKE PLACE
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// // gracefully shutdown the server, waiting max 30 seconds for current operations to complete

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
