package config_test

import (
	"testing"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/stretchr/testify/assert"
)

var (
	Config config.Config
)

func TestLoadConfig(t *testing.T) {
	envFilePath := "../../cmd/.env"

	Config, err := config.LoadConfig(envFilePath)

	assert.Nil(t, err)
	assert.NotNil(t, Config.Server.Address)
	assert.NotNil(t, Config.Server.ClientAddress)
	assert.NotNil(t, Config.Server.MaxFileSizeInMb)
	assert.NotNil(t, Config.FoodData.Address)
	assert.NotNil(t, Config.FoodData.ApiKey)
	assert.NotNil(t, Config.Database.Username)
	assert.NotNil(t, Config.Database.Password)
	assert.NotNil(t, Config.Database.Address)
	assert.NotNil(t, Config.QrDecoder.Address)
}
