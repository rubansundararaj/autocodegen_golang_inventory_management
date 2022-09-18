package main

import (
	"fmt"
	"log"
	"inventory_management/router"
	"net/http"
)

func main() {
	fmt.Println("Starting inventory_management API...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000....")
}
