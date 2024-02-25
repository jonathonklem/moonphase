package main

// import the communication package

import (
	// import the communication package
	"moonalert/communication"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	communication.ProcessStates();
}