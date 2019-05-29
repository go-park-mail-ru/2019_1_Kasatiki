package game_logic

import "testing"

func TestBullet_Run(t *testing.T) {
	b := &Bullet{
		Angle: 3.14,
	}
	b.Object = &DynamycObject{
		X:        10,
		Y:        10,
		Velocity: 10,
	}
	b.Run()
	if b.Object.X == 10 {
		t.Errorf("Failed bullet flying")
	}
}
