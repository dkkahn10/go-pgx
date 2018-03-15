package util

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

//Connect to Windows environment
func Connect(applicationName string) (conn *pgx.Conn) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		User:              os.Getenv("POSTGRES_USER"),
		Password:          os.Getenv("POSTGRES_PASSWORD"),
		Host:              os.Getenv("POSTGRES_HOST"),
		Port:              5432,
		Database:          "test-rating",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
		os.Exit(1)
	}
	return conn
}
