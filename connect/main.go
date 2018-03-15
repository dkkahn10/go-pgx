package main

import (
	"fmt"
	"os"

	"github.com/dkkahn10/go-pgx"
)

func main() {
	conn := util.Connect("Creating table in Windows postgreSQL")
	defer conn.Close()

	_, err := conn.Exec(`
		CREATE TABLE cards (
			id integer NOT NULL,
			board_id integer NOT NULL,
			data jsonb
		);
	`)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create users table: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Connection worked!\n")
}
