package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

type DBConn struct {
	Name     string
	User     string
	Port     int
	Password string
}

func loadDB() (*DBConn, error) {
	p := DBConn{
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}

	portString := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to int: %v", err)
	}

	p.Port = port

	if p.Name == "" || p.User == "" || p.Password == "" {
		return nil, fmt.Errorf("cannot have null fields! p.Name = %v, p.User = %v, p.Password = %v",
			p.Name, p.User, p.Password)
	}

	return &p, nil
}

func ConnectDB() (*sql.DB, error) {
	connVals, err := loadDB()
	if err != nil {
		return nil, fmt.Errorf("failed to load db conn vals: %v", err)
	}

	connString := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable",
		connVals.User, connVals.Password, connVals.Name, connVals.Port)

	dbConn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	defer db.Close()

	db = dbConn

	return db, nil
}

func DBInit() error {
	filepath := "schema.sql"

	file, err := os.ReadFile(filepath)
	if err != nil {
		wd, _ := os.Getwd()
		return fmt.Errorf("failed to read file: filepath: %v, currDir: %v, content: %v",
			filepath, wd, file)
	}

	query := (string(file))

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to run db exec. query: %v, err: %v", query, err)
	}

	return nil
}
