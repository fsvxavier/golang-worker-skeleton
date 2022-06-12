package enviroment_test

import (
	"os"
	"testing"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/enviroment"
	"github.com/stretchr/testify/assert"
)

func TestEnvConfig(t *testing.T) {
	os.Setenv("TEST_ENV_INT", "1")
	os.Setenv("ENV", "production")

	env := enviroment.NewEnviroment()

	ret, err := env.GetTag("TEST_ENV_INT")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvString(t *testing.T) {
	os.Setenv("ENV", "production")

	env := enviroment.NewEnviroment()
	env.SetFileConfig("./mock/env.json")

	ret, err := env.GetTag("TEST_ENV_STRING")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvEmpty(t *testing.T) {
	os.Setenv("ENV", "production")

	env := enviroment.NewEnviroment()
	env.SetFileConfig("./mock/env.json")

	ret, err := env.GetTag("TEST_ENV_EMPTY")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvBool(t *testing.T) {

	os.Setenv("ENV", "production")

	env := enviroment.NewEnviroment()
	env.SetFileConfig("./mock/env.json")

	ret, err := env.GetTag("TEST_ENV_BOOL")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvInt(t *testing.T) {
	os.Setenv("ENV", "production")

	env := enviroment.NewEnviroment()
	env.SetFileConfig("./mock/env.json")

	ret, err := env.GetTag("TEST_ENV_INT")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}
