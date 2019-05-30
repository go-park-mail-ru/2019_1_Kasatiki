package lobby

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/connections"
	"testing"
)

func TestNewLobby(t *testing.T) {
	l := NewLobby()
	if l != nil {
		t.Errorf("Failed lobby creating")
	}
}

func TestLobby_AddPlayer(t *testing.T) {
	c := &connections.UserConnection{}
	l := NewLobby()
	err := l.AddPlayer(c)
	if err != nil {
		t.Errorf("Failed add player")
	}

}
