package storage

import (
	"database/sql"
   "github.com/anthdm/Gobank/types"
	_ "github.com/lib/pq")

type Storage interface {

	CreateAccount(*types.Account) error
    DeleteAccount(int) error
	UpdateAcoount(*types.Account) error
	GetAcocuntByID(int)(*types.Account,error)     
}
type PostgresStore struct {
		db *sql.DB}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=some-postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil}
func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()}	
func (s *PostgresStore) CreateAccountTable() error {
querry := `create table if not exist  accounts(
id SERIAL PRIMARY KEY, 
firstname varchar(50), 
lastname varchar(50), 
number serial,
balance serial,
created_at TIMESTAMP)`
_, err := s.db.Exec(querry)
return err}

 
func (s *PostgresStore) CreateAccount(a *types.Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {return nil

}
func (s *PostgresStore) UpdateAccount(*types.Account) error {return nil}
func (s *PostgresStore) GetAccountByID(id int) (*types.Account, error) {return nil, nil}
func (s *PostgresStore) GetAccounts() (*types.Account, error) {return nil, nil}
