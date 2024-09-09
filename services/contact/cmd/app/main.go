package main

import (
	"Contact_service/pkg/store/clickhouse"
	"fmt"
)

func main() {
	conn, err := clickhouse.New(clickhouse.Settings{})
	if err != nil {
		panic(err)
	}

	defer conn.Conn.Close()
	fmt.Println("Connection to ClickHouse established")
}
