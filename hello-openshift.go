package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	helloEnv := os.Getenv("HELLO")

	if helloEnv == "" {
		helloEnv = "tsdf hello"
	}

	arg := "openshift"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	log.Printf("%s %s\n", helloEnv, arg)

	fmt.Fprintf(w, "%s %s\n", helloEnv, arg)
}

func main() {
	port := "8080"
	
	var portEnv string
	portEnv = os.Getenv("PORT")
	
	if portEnv != "" {
		port = portEnv
	}
	
	http.HandleFunc("/", hello)
	log.Println("Starting server on 0.0.0.0:"+port)
	http.ListenAndServe(":"+port, nil)
}
