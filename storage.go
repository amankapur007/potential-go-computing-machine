package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host=172.29.224.1 user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number int,
		balance int,
		created_at timestamp
	)`

	_, err := s.db.Query(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account (first_name, last_name, number, balance, created_at)
	values ($1, $2, $3,$4,$5)`

	res, err := s.db.Query(query,
		acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}
	fmt.Printf("Res %v", res)
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query(`delete from account where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := `select * from account where id = $1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanToAccount(rows)

	}
	return nil, fmt.Errorf("account is not found for id %d", id)
}
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	query := `select * from account`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := ScanToAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func ScanToAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)
	return account, err
}
