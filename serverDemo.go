/*
Example http server.
Based on package documentation from golang.org.
Code is licensed under a BSD license.
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve content from folder on filesystem
	root := http.FileServer(http.Dir("hello"))
	// Register var root as the handler of requests matching "/"
	http.Handle("/", root)
	// Print a message to stdout when server starts
	log.Println("Server is running. Is there anybody out there?")

	/*
		Listen on specified port for requests. 2nd @arg is the handler for
		incoming requests. Passing nil causes the function to use the
		handler we already defined with http.Handle.
	*/
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
