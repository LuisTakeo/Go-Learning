package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Takeo", "Chicão", "Charlinha"}

	// Get a greeting message and print it.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(messages[names[1]])
	// fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}

// Run the following command to install the module:
// go mod init example.com/greetings -> create a go.mod file
// go mod tidy -> add the module's dependencies to the go.mod file
// go run . -> run main.go
// go build -> build the module into an executable file
// go install -> install the module as a binary in the workspace's bin directory

// Para rodar o programa:
// go mod init example.com/greetings -> cria um arquivo go.mod
// go mod tidy -> adiciona as dependencias do modulo no arquivo go.mod
// go run -> roda o arquivo main
// go build -> cria um executavel
// go install -> cria um executavel e instala no workspace/bin
