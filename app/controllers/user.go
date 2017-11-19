package controllers

import (
	"fmt"

	"github.com/gocraft/web"
)

func Ping(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "pong")
}
