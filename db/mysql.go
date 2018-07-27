package db

import (
	"database/sql"
	"fmt"
	"go-graphql/config"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	conn *sql.DB
}

var DB = Database{

}

type MySQLConfig struct {
	Username   string
	Password   string
	Host       string
	Port       int
	UnixSocket string
}

func newDatabase() (Database, error) {

	conn, err := sql.Open("mysql", config.MysqlConnectURL)

	if err != nil {
		return DB, fmt.Errorf("mysql: could not get a connection: %v", err)
	}

	if err := conn.Ping(); err != nil {
		conn.Close()
		return DB, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	DB.conn = conn

	return DB, err
}

func (db Database) Close() {
	db.conn.Close()
}

func (db Database) Query(query string, args interface{}) (*sql.Rows, error) {
	return DB.conn.Query(query, args)
}

func (db Database) Prepare(query string) (*sql.Stmt, error) {
	return DB.conn.Prepare(query)
}

func InitDatabase() (Database) {
	DB, err := newDatabase()
	if err != nil {
		fmt.Errorf("error connect to database: %v", err)

	}

	return DB
}

func (db Database) Insert(query string, args ...interface{}) (int64, error) {

	stmt, _ := DB.conn.Prepare(query)

	r, err := stmt.Exec(args...)
	if err != nil {
		return 0, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return 0, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}

	lastInsertID, err := r.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	return lastInsertID, nil

}

func (db Database) Update(query string, args ...interface{}) (int64, error) {


	fmt.Println("Update" ,query, args)

	stmt, _ := DB.conn.Prepare(query)

	r, err := stmt.Exec(args...)
	if err != nil {
		return 0, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return 0, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}

	lastInsertID, err := r.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}

	return lastInsertID, nil

}