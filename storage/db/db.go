package db

import (
	"github.com/MiguelMoll/joycast/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	client *gorm.DB
}

func New(conn string) (*DB, error) {
	client, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return &DB{client: client}, nil
}

func (d *DB) Close() error {
	return d.client.Close()
}

func (d *DB) CreateUser(user *types.User) error {
	du, err := dbUser(user)
	if err != nil {
		return err
	}
	return d.client.Create(du).Error
}

func (d *DB) GetUser(id uint) (*types.User, error) {
	var user User
	err := d.client.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	su, err := storeUser(&user)

	return su, nil
}

func (d *DB) SaveUser(user *types.User) error {
	du, err := dbUser(user)
	if err != nil {
		return err
	}

	return d.client.Save(du).Error
}
