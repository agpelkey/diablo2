package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Interface to hold DB functions
type PostgresRepo interface {
}

type PostgresDB struct {
	db *sql.DB
}

// function to create new DB
func (p *PostgresDB) NewPostgresDB() (*PostgresDB, error) {
	// we will be running postgres as a docker image for this project
	// the following logic will peratin to accessing that container

	connstr := "user=postgres dbname=diablo2 password=diablo2 sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

// function to initialize DB
func (p *PostgresDB) Init() error {
	// logic to initialize create table function
	return p.CreateUserAccountTable()
}

// function to create user account table
func (p *PostgresDB) CreateUserAccountTable() error {
	// sql logic to create table in DB
	query := `CREATE TABLE IF NOT EXISTS users (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		username varchar(50),
		email varchar(50),
		encrypted_password varchar(100),
		created_at timestamp
	);`

	_, err := p.db.Exec(query)
	return err
}

// function to get account by ID

// function to create account

// function to delete account
