package main

import (
	"context"
	"errors"
	"fmt"
    "fourth/pkg/consts"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
    log.Println(consts.SrvStrd)

	ListeningStart()
}

func ListeningStart() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
            ctx = context.WithValue(ctx, consts.KeySvrAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	first := r.URL.Query().Get("first")
	second := r.URL.Query().Get("second")

	log.Printf("%s: got / request. first=%s, second=%s\n",
        ctx.Value(consts.KeySvrAddr),
		first,
		second)

	io.WriteString(w, "This is my website!\n")
}
