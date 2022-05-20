package main

import (
	"fmt" 
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!\n")
    fmt.Println("Endpoint hit: homePage")
}

func hello(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "hello %s!\n", os.Getenv("NAME"))
    fmt.Println("Endpoint hit: hello")
}

func health(w http.ResponseWriter, r *http.Request){
    var health bool = true
	fmt.Fprintf(w, "health: %v!\n", health)
    fmt.Println("Endpoint hit: health")
}

type Counter struct {
    counter int
}

var serverCounter Counter = Counter {counter: 0}

func count(w http.ResponseWriter, r *http.Request){
    serverCounter.counter += 1
	fmt.Fprintf(w, "counter: %v!\n", serverCounter.counter)
    fmt.Println("Endpoint hit: count")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/health", health)
    http.HandleFunc("/count", count)
   
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
