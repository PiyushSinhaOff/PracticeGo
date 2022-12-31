package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func run(ch chan struct{}) {
	// uncomment this to avoid server crash
	defer func() {
		close(ch)
		r := recover()
		if r != nil {
			log.Println("panic happened", r)
		}
	}()

	panic("helloWorld")
	//fmt.Println("Hello World123!")
	return
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		log.Println("Bye", r)
	}()
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Endpoint Hit: helloWorld")
	ch := make(chan struct{})
	go run(ch)
	<-ch
}

func abc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABC!")
	fmt.Println("Endpoint Hit: abc")
}

func handleRequests() {
	defer func() {
		r := recover()
		log.Println(r)
	}()
	r := mux.NewRouter()
	r.HandleFunc("/panic", helloWorld)
	r.HandleFunc("/abc", abc)
	log.Fatal(http.ListenAndServe(":8097", r))
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			log.Println("panic happened in main", r)
		}
	}()
	panic("helloWorljake")
	handleRequests()
}
