package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello !!!First time on Browser via Golang!!!")
	_, err := writer.Write(message)
	check(err)
}

func main() {
	http.HandleFunc("/hello", viewHandler)            //tells the app to call viewHandler whenever a request for a URL ending in /hello is received.
	err := http.ListenAndServe("localhost:8080", nil) //The nil value in the second argument just means that requests will be handled using functions set up via HandleFunc
	/* if you want to learn about alternate ways to handle requests,
	look up the documentation for the “ListenAndServe” function,
	the “Handler” interface, and the “ServeMux” type from the “http” package.)*/
	log.Fatal(err) //check() is not invoked as ListenAndServe will always return with error
}
