package tests

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"path/filepath"
	"time"
)

var (
	dbName     = "chaindexing_test"
	dbUser     = "test_user"
	dbPassword = "test_password"
)

func NewPostgresDB(ctx context.Context) *pg.PostgresContainer {
	postgresContainer, err := pg.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:14.5-alpine"),
		pg.WithInitScripts(filepath.Join("testdata", "init-test-db.sh")),
		//postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		pg.WithDatabase(dbName),
		pg.WithUsername(dbUser),
		pg.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	return postgresContainer
}

func GormPostgresInstance(connString string) (db *gorm.DB) {
	db, err := gorm.Open(
		postgres.Open(connString),
	)

	if err != nil {
		panic(err)
	}

	return
}
