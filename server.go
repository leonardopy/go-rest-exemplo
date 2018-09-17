package main

import (
	"net/http"
	"planeta/routes"
)

func main() {
	http.ListenAndServeTLS("0.0.0.0:3010", "/home/ubuntu/certs/fullchain.pem", "/home/ubuntu/certs/privkey.pem", routes.R)
}
