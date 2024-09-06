package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/Shanmuganthan/go-lang-mongo/router"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func LoadEnv() {

	envType := flag.String("env", "development", "Specify the environment type: development, prod, etc.")

	flag.Parse()

	envFile := ".env"

	switch *envType {
	case "development":
		envFile = ".env.development"
	case "production":
		envFile = ".env.production"
	default:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)

	if err != nil {
		fmt.Println("Loading Env Issue", err)
	}

	fmt.Println("Env File Loaded", envFile)

}

func main() {
	file, fileerr := os.OpenFile("D:\\NATHAN\\Go Projects\\go-lang-mongo\\app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileerr != nil {
		log.Printf("Failed to open log file: %v", fileerr)
	}
	defer file.Close()

	log.Println("File created")

	r := router.Router()

	log.SetOutput(file)

	LoadEnv()

	err := http.ListenAndServe(":4000", r)

	log.Println("Aplication Initalization")

	if err != nil {
		log.Fatal(fmt.Printf("Server creation Failed ", err))
	}

	log.Println("Server Started at 4000")

	defer func() {
		log.Println("Application exited")
	}()

}
