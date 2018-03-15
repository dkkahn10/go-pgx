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

	var tags string
	err := conn.QueryRow(`
		SELECT jsonb_array_elements_text(data->'tags') AS tag
		FROM cards 
		WHERE id=2
	`).Scan(&tags)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Fprint(os.Stderr, "No rows found")
		} else {
			fmt.Fprintf(os.Stderr, "Unsuccessful query: %v\n", err)
		}
		os.Exit(1)
	}
	fmt.Printf("Found some tags %s\n", tags)
}
