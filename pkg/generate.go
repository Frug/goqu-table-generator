package pkg

import (
	"database/sql"
	"fmt"

	"github.com/Frug/goqu-table-generator/config"
	_ "github.com/lib/pq"
)

func Generate(conf config.Config) (err error) {
	fmt.Printf("Generating scaffolding for database %s.\n", conf.DBName)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPass, conf.DBName,
	)

	con, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("connecting to postgres: %w", err)
	}

	if err := con.Ping(); err != nil {
		return fmt.Errorf("checking database connection: %w", err)
	}

	return nil
}
