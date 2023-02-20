package pkg

import "testing"

func TestReadEnv(t *testing.T) {
    env, _ := ReadEnv()
    t.Logf("%+v", env)
}
