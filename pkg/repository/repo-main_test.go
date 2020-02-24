package repository

import (
	"fmt"
	"os"
	"testing"

	"pgxs.io/chassis"
	"pgxs.io/chassis/config"
	"pgxs.io/qurl/pkg/migrations"

	cfg "pgxs.io/qurl/pkg/config"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	cfg.ResetEnvKey()
	config.LoadFromEnvFile()
	err := migrations.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	QurlRepositoryInstance()
}

func teardown() {
	chassis.CloseDB()
}
