package main

import (
    "errors"
    "fourth/pkg/constants"
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    log.Println(constants.LogServerStarted)
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)

    ListeningStart()
}

func ListeningStart() {
    if err := http.ListenAndServe("", nil);
    errors.Is(err, http.ErrServerClosed) {
        log.Printf("server closed\n")
    } else if err != nil {
        log.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
	io.WriteString(w, "Not-using URL!\nPlease, type correct.\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /hello request\n")
    io.WriteString(w, "Hello, HTTP!\n")
}