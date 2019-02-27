package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/oauth2"

	"github.com/MiguelMoll/joycast/storage"
)

type DB struct {
	db *gorm.DB
}

func New(conn string) (*DB, error) {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Product{})
	return &DB{db: db}, nil
}

func (d *DB) SaveToken(user *storage.User, token *oauth2.Token) error {
	return nil
}

func (d *DB) GetToken(user *storage.User) (*oauth2.Token, error) {
	return nil, nil
}
