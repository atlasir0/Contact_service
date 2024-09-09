package main

import (
	"Contact_service/pkg/store/postgres"
	"fmt"
)

func main() {
	conn , err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()
	fmt.Println(conn.Pool.Stat())
	fmt.Println("Hello")
}
