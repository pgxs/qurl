package service

import (
	"fmt"
	"os"
	"testing"

	"pgxs.io/chassis"
	"pgxs.io/chassis/config"

	cfg "pgxs.io/qurl/pkg/config"
	"pgxs.io/qurl/pkg/migrations"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	cfg.ResetEnvKey()
	config.LoadFromEnvFile()
	if err := migrations.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	QUrlServiceInstance()
}
func tearDown() {
	chassis.CloseDB()
}
