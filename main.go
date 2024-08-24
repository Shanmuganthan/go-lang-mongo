package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shanmuganthan/go-lang-mongo/router"
)

func main() {

	fmt.Println("Aplication Initalization")

	r := router.Router()

	fmt.Println("Aplication Initalization")
	err := http.ListenAndServe(":4000", r)

	if err != nil {
		log.Fatal("Server creation Failed", err)
	}

	fmt.Println("Server Started at 4000")

}
