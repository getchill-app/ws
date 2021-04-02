package main

import (
	"log"
	"os"

	"github.com/getchill-app/ws/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	sk := os.Getenv("SECRET_KEY")
	tk := os.Getenv("TOKEN_KEY")
	log.Fatal(server.ListenAndServe(":8080", "ws://localhost:8080/ws", sk, tk, nil))
}
