package main

import (
	"fmt"

	"github.com/dkkahn10/go-pgx"
)

func main() {
	conn := util.Connect("connect test")
	defer conn.Close()
	fmt.Printf("Connection worked!\n")
}
