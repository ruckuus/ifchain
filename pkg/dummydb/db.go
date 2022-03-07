package dummydb

import (
	"fmt"
	"log"
	"os"
)

type data map[string]interface{}

type DB struct {
	ns  string
	log *log.Logger

	database data
}

func NewDB(ns string) *DB {
	return &DB{
		ns:       ns,
		log:      log.New(os.Stdout, fmt.Sprintf("db_log: %s: ", ns), log.Lshortfile|log.Ltime|log.Lmicroseconds),
		database: make(map[string]interface{}),
	}
}

func (c *DB) Set(key string, val interface{}) error {
	c.log.Printf("Saving object: %v", val)
	c.database[key] = val
	return nil
}

func (c *DB) Get(key string) (interface{}, error) {
	if val, ok := c.database[key]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("not found: %s", key)
}
