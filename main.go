package main

import (
	"fmt"
	"os"
	"test-server/connection"
)

func main() {
	listener, err := connection.OpenServer(os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
}
