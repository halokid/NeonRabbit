package pkg

import (
	"path/filepath"
	"testing"
)

func TestReadEnv(t *testing.T) {
	// graceful access the file path
	envFile := "../config.yaml"
	absFilePath, err := filepath.Abs(envFile)
	t.Logf("file path -->>> %+v, err -->>> %+v", absFilePath, err)

	env, _ := ReadEnv()
	t.Logf("%+v", env)
}
