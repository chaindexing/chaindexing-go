package tests

import (
	"context"
	"github.com/chaindexing/chaindexing-go"
	"github.com/chaindexing/chaindexing-go/tests/testdata"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"reflect"
	"testing"
)

func TestCreateMany(t *testing.T) {
	ctx := context.Background()
	//ctrl := gomock.NewController(t)
	postgresContainer := NewPostgresDB(ctx)
	connStr, _ := postgresContainer.ConnectionString(ctx, "sslmode=disable", "application_name=chaindexing")

	// terminate container
	defer func(postgresContainer *postgres.PostgresContainer, ctx context.Context) {
		err := postgresContainer.Terminate(ctx)
		if err != nil {
			panic(err)
		}
	}(postgresContainer, ctx)

	db := chaindexing.Database{DB: GormPostgresInstance(connStr)}

	cAddressRepo := chaindexing.NewContractAddressRepository(&db)
	amountOfCAddressesToCreate := 10

	// TABLE TESTS
	var tests = []struct {
		name  string
		input any
		want  any
	}{
		// the table itself
		{"C-Addresses should be created without error", amountOfCAddressesToCreate, nil},
		{"DB Table count should correspond", nil, amountOfCAddressesToCreate},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		err := cAddressRepo.CreateMany(testdata.GenerateContractAddresses(int(reflect.ValueOf(tests[0].input).Int()))...)

		if err != tests[0].want {
			t.Errorf("Result was incorrect, got: %v, want: %v", err, tests[0].want)
		}
	})

	t.Run(tests[1].name, func(t *testing.T) {
		count := cAddressRepo.Count()

		if count != int64(reflect.ValueOf(tests[1].want).Int()) {
			t.Errorf("Result was incorrect, got: %v, want: %v", count, tests[1].want)
		}
	})

}
