package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildLogger(t *testing.T) {
	t.Run("test BuildLogger default env success", func(t *testing.T) {
		builtLogger := BuildLogger()
		assert.NotNil(t, builtLogger)
	})

	t.Run("test BuildLogger production env success", func(t *testing.T) {
		os.Setenv("AppEnv", "production")
		builtLogger := BuildLogger()
		assert.NotNil(t, builtLogger)
	})
}
