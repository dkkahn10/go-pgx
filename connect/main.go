package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"

	"github.com/dkkahn10/go-pgx"
)

func main() {
	conn := util.Connect("Creating table in Windows postgreSQL")
	defer conn.Close()

	var tag string
	rows, err := conn.Query(`
		SELECT jsonb_array_elements_text(data->'tags') AS tag
		FROM cards 
		WHERE id=2
	`)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Fprint(os.Stderr, "No rows found")
		} else {
			fmt.Fprintf(os.Stderr, "Unsuccessful query: %v\n", err)
		}
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tag)
		fmt.Printf("Tag: %s\n", tag)
	}
}
