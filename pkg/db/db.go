package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DB_USERNAME  = "DB_USERNAME"
	DB_USER_PASS = "DB_USER_PASS"

	DB_NAME = "DB_NAME"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"

	DB_AUTH_SOURCE = "DB_AUTH_SOURCE"
)

const cancelDbQueryTimeoutSc = 4 * time.Second

type DB struct {
	Client *mongo.Client
}

type db struct {
	name string
	host string
	port string

	userName string
	userPass string

	authSource string
}

func (db *DB) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), cancelDbQueryTimeoutSc)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(newDBUrl()))
	if err != nil {
		log.Fatal(err)
	}

	db.Client = client
}

func (db *DB) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), cancelDbQueryTimeoutSc)
	defer cancel()

	err := db.Client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), cancelDbQueryTimeoutSc)
	defer cancel()

	err := db.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}

func getDbInfoFromEnv() *db {
	return &db{
		name:       os.Getenv(DB_NAME),
		host:       os.Getenv(DB_HOST),
		port:       os.Getenv(DB_PORT),
		userName:   os.Getenv(DB_USERNAME),
		userPass:   os.Getenv(DB_USER_PASS),
		authSource: os.Getenv(DB_AUTH_SOURCE),
	}
}

func newDBUrl() string {
	dbInfo := getDbInfoFromEnv()

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		dbInfo.userName, dbInfo.userPass, dbInfo.host, dbInfo.port, dbInfo.name, dbInfo.authSource,
	)
}
