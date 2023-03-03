package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const dbtimeout = time.Minute * 3

// Interface to hold DB functions
type PostgresRepo interface {
	CreateUserAccount(acc *Account) error
	GetUserByID(id int) (*Account, error)
}

type PostgresDB struct {
	db *sql.DB
}

// function to create new DB
func NewPostgresDB() (*PostgresDB, error) {
	// we will be running postgres as a docker image for this project
	// the following logic will peratin to accessing that container

	connstr := "user=postgres dbname=postgres password=d2 sslmode=disable"
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
func (p *PostgresDB) GetUserByID(id int) (*Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `SELECT * FROM users where id = $1`
	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		ScanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

// function to create account

// function to delete account

// function to create user account in DB
func (p *PostgresDB) CreateUserAccount(acc *Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `INSERT INTO users (
		first_name, last_name, username, email, encrypted_password, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := p.db.QueryContext(ctx, query,
		acc.FirstName,
		acc.LastName,
		acc.Username,
		acc.Email,
		acc.EncryptedPassword,
		acc.CreatedAt)

	if err != nil {
		fmt.Println(err)
	}

	return nil

}

func ScanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Username,
		&account.Email,
		&account.EncryptedPassword,
		&account.CreatedAt)

	return account, err
}
