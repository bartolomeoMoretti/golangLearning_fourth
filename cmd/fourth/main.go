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

	ListeningStart()
}

func ListeningStart() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	if err := http.ListenAndServe("", mux); errors.Is(err, http.ErrServerClosed) {
		log.Printf(constants.LogServerClosed)
	} else if err != nil {
		log.Println(constants.LogServerErrSt, err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf(constants.LogRootAddr)
	io.WriteString(w, constants.RootAddr)
}
func getHello(w http.ResponseWriter, r *http.Request) {
	log.Printf(constants.LogHelloAddr)
	io.WriteString(w, constants.HelloAddr)
}

func (receiver) name() {

}
