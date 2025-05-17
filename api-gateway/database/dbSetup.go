package database

import (
	"database/sql"
	"fmt"
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
	Host     string
}

func loadDB() (*DBConn, error) {
	p := DBConn{
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
	}

	portString := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to int: %v", err)
	}

	p.Port = port

	if p.Name == "" || p.User == "" || p.Password == "" || p.Host == "" {
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

	connString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", connVals.Host,
		connVals.User, connVals.Password, connVals.Name, connVals.Port)

	dbConn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	db = dbConn

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db. err %v: ", err)
	}

	return db, nil
}

func DBInit() error {
	tablesExist, err := tablesExists(db)
	if err != nil {
		return fmt.Errorf("Failed to check if tables existed, err: ", err)
	}

	if tablesExist {
		return nil
	}

	filepath := "database/schema.sql"

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

	fmt.Print("Successfully connected to db")

	return nil
}

func tablesExists(db *sql.DB) (bool, error) {
	query := `SELECT COUNT(*) FROM information_schema.tables`

	result := db.QueryRow(query)
	var count int

	err := result.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to scan rows into var, err: %v", err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
