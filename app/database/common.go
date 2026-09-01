package database

import (
	"errors"
	"log"
	"time"
)

const connectRetries = 5

func Connect(config *Config, retryInterval time.Duration) (*DB, error) {
	for retry := 0; retry < connectRetries; retry++ {
		db, err := New(config)
		if err == nil {
			if err := db.Ping(); err != nil {
				return nil, err
			}
			return db, nil
		}
		log.Println(err)
		log.Printf("unable to connect to database: sleeping...\n")
		time.Sleep(retryInterval)
	}
	return nil, errors.New("could not connect to database")
}
