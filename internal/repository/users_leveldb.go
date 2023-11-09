package repository

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"log"
)

type UsersLevelDB struct {
	db *leveldb.DB
}

func NewUsersLevelDB(dbPath string) *UsersLevelDB {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Panicf("while opening leveldb from file: %v", err)
	}
	return &UsersLevelDB{db}
}

func (u *UsersLevelDB) Create(login, hashedPassword string) error {
	err := u.db.Put([]byte(login), []byte(hashedPassword), &opt.WriteOptions{NoWriteMerge: true})
	if err != nil {
		return fmt.Errorf("while putting in level db: %w", err)
	}
	return nil
}

func (u *UsersLevelDB) GetStoredPass(login string) (string, error) {
	pass, err := u.db.Get([]byte(login), nil)
	if err != nil {
		return "", fmt.Errorf("while getting pass from leveldb: %v", err)
	}
	return string(pass), err
}
