package migrations

import (
	"os"
	"pgxs.io/chassis/config"
	cfg "pgxs.io/qurl/pkg/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(-1)
	}
	exitVal := m.Run()

	//tearDown
	os.Exit(exitVal)
}
func setup() error {
	cfg.ResetEnvKey()
	config.LoadFromEnvFile()
	return nil
}
func Test_ImportData(t *testing.T) {
	err := Run()
	assert.NoError(t, err)
	if !chassis.EnvIsProd() {
		count := 0
		chassis.DB().Table("qurls").Count(&count)
		assert.NotEmpty(t, count)
		assert.Equal(t, 1, count)
	}
}
