package main

import (
	"fourth/internal/app/endpoint"
	"fourth/pkg/consts"
	"log"
)

func main() {
	log.Println(consts.SrvStrd)

	endpoint.New()
}
