package dbhandler

import "testing"

func TestRandStr(t *testing.T) {
	random := RandStr(10)
	if len(random) != 10 {
		t.Errorf("Wrong len")
	}
}
