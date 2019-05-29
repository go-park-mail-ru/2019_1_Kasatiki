package lobby

import "testing"

func TestNewLobby(t *testing.T) {
	l := NewLobby()
	if l != nil {
		t.Errorf("Failed lobby creating")
	}
}
