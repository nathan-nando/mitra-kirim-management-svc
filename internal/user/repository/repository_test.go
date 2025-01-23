package repository

import (
	"mitra-kirim-be-mgmt/config"
	"mitra-kirim-be-mgmt/http/rest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatusBykey(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Error("Error loading config", err)
		return
	}
	logger := rest.NewLogger()
	db := config.NewDatabase(cfg, logger)
	repo := User{Db: db}
	result, err := repo.FindById(2)
	t.Log(result)
	assert.NoError(t, err)
	//assert.Equal(t, mockStatus, status)
}
