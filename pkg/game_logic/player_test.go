package game_logic

import (
	"testing"
)

func TestPlayer_CreateDefaultWeapon(t *testing.T) {
	p := &Player{}
	p.Spawn(1, 2, 50, 50)
	p.CreateDefaultWeapon()
	if p.Weapon.Name != "Deagle" {
		t.Errorf("Failed player create default weapon")
	}
}

func TestPlayer_Spawn(t *testing.T) {
	p := &Player{}
	//p.CreateDefaultWeapon()
	p.Spawn(1, 2, 50, 50)
	if p.Object.Y != 2 {
		t.Errorf("Failed player spawn")
	}
}

func TestPlayer_Shot(t *testing.T) {
	p := &Player{}
	p.Spawn(1, 2, 50, 50)
	p.CreateDefaultWeapon()
	b := p.Shot(3.14)
	if b == nil {
		t.Errorf("Failed player shot")
	}

}

func TestPlayer_SetAngular(t *testing.T) {

	p := &Player{}
	p.SetAngular(3.14)
	if p.Angular != 3.14 {
		t.Errorf("Failed player set angular")
	}

}

func TestPlayer_SetNickname(t *testing.T) {

	p := &Player{}
	p.SetNickname("Nick")
	if p.Nickname != "Nick" {
		t.Errorf("Failed player set nickname")
	}

}
