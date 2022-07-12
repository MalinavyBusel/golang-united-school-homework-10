package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	Router := mux.NewRouter()

	Router.HandleFunc("/name/{PARAM}",
		handleNameGet).Methods(http.MethodGet)
	Router.HandleFunc("/bad",
		handleBadConn).Methods(http.MethodGet)
	Router.HandleFunc("/data",
		handleDataPost).Methods(http.MethodPost)
	Router.HandleFunc("/headers",
		handleHeadersPost).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), Router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func handleNameGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["PARAM"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func handleBadConn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError) // 500
}

func handleDataPost(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", data)))
	w.WriteHeader(http.StatusOK) // 200
}

func handleHeadersPost(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	w.Header().Set("a+b", strconv.Itoa(a+b))
	w.WriteHeader(http.StatusOK) // 200
}
