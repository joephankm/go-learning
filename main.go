package main

import (
	"net/http"

	"github.com/hello-world/app/route"
)

func main() {
	http.ListenAndServe("localhost:7000", route.Router())
}
