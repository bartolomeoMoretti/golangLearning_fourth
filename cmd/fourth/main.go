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

func (receiver ) name()  {
    
}