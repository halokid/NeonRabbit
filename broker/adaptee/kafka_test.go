package adaptee

import "testing"

func TestKafka_Pub(t *testing.T) {
    k := NewKafka()
    k.Pub("neon_rabbit", "test message")
}
