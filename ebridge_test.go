package ebridge

import (
	"testing"
)

func TestLongPVGet(t *testing.T) {
	defer teardown()
	d := edemo()
	t.Fatalf("d : %d", d)
}
