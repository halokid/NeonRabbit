package log

import "testing"

func TestGetAll(t *testing.T) {
  logs := GetAll()
  t.Logf("logs -->>> %+v", logs)
}
