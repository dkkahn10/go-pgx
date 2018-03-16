package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dkkahn10/go-pgx"
	"github.com/jackc/pgx"
)

type data struct {
	Task    string   `json:"Task"`
	Finshed bool     `json:"finished"`
	Tags    []string `json:"tags"`
}

func main() {
	conn := util.Connect("Creating table in Windows postgreSQL")
	defer conn.Close()

	_, indexErr := conn.Exec(`
		CREATE INDEX idxtag ON cards((data->>'tags'));
	`)

	if indexErr != nil {
		fmt.Fprintf(os.Stderr, "Unsuccessful indexing: %v\n", indexErr)
	}

	var count int
	number, countErr := conn.Query(`
		SELECT COUNT(*) FROM cards;
	`)

	if countErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to return row count: %v\n", countErr)
	}

	for number.Next() {
		number.Scan(&count)
		fmt.Printf("Row count is: %v\n", count)
	}

	_, insErr := conn.Exec(`
			INSERT INTO cards VALUES (
				$1, 6, '{"Task": "Mow the lawn", "tags": ["Lawn", "Grass", "Mow"], "finished": false}'
			);
	`, count+1)

	if insErr != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert new row: %v\n", insErr)
	}

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

	// items, itemErr := conn.Query(`
	// 	SELECT id FROM cards WHERE data ->> 'items' IS NOT NULL;
	// 	`)

	// if itemErr != nil {
	// 	fmt.Printf("There was an error: %v\n", itemErr)
	// }

	// for items.Next() {
	// 	var id int32
	// 	items.Scan(&id)
	// 	fmt.Printf("The item is: %d\n", id)
	// }

	// defer items.Close()

	items, itemErr := conn.Query(`
		SELECT * FROM cards WHERE data -> 'finished' = 'true';	
		`)

	if itemErr != nil {
		fmt.Printf("There was an error: %v\n", itemErr)
	}

	var strTasks []data
	for items.Next() {
		var id int32
		var bid int32
		var array []byte
		scanErr := items.Scan(&id, &bid, &array)
		if err != nil {
			fmt.Println(scanErr)
		}

		result := data{}
		json.Unmarshal(array, &result)
		strTasks = append(strTasks, result)

	}

	fmt.Printf("%v", strTasks)

	defer items.Close()
}
