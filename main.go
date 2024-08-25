package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/Shanmuganthan/go-lang-mongo/router"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	file, fileerr := os.OpenFile("D:\\NATHAN\\Go Projects\\go-lang-mongo\\app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileerr != nil {
		log.Printf("Failed to open log file: %v", fileerr)
	}
	defer file.Close()

	log.Println("File created")

	r := router.Router()

	err := http.ListenAndServe(":4000", r)

	log.SetOutput(file)

	log.Println("Aplication Initalization")
	if err != nil {
		log.Fatal(fmt.Printf("Server creation Failed ", err))
	}

	log.Println("Server Started at 4000")

	defer func() {
		log.Println("Function is cloisong")
	}()

}
