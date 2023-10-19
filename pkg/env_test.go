package pkg

import (
  "testing"
)

func TestReadEnv(t *testing.T) {
  // graceful access the file path
  //envFile := "../config.yaml"
  //absFilePath, err := filepath.Abs(envFile)
  //t.Logf("file path -->>> %+v, err -->>> %+v", absFilePath, err)
  //

  pkgxTmp := Pkg{}
  env, _ := ReadEnv(&pkgxTmp)
  t.Logf("%+v", env)
}
