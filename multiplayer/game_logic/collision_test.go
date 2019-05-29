package game_logic

import "testing"

func TestIsCollision(t *testing.T) {
	obj1 := &DynamycObject{
		X:     10,
		Y:     10,
		Xsize: 10,
		Ysize: 10,
	}
	obj2 := &DynamycObject{
		X:     15,
		Y:     15,
		Xsize: 10,
		Ysize: 10,
	}
	IsCollision(obj1, obj2)
	if !IsCollision(obj1, obj2) {
		t.Errorf("Failed test collision")
	}
	obj3 := &DynamycObject{
		X:     105,
		Y:     105,
		Xsize: 10,
		Ysize: 10,
	}
	if IsCollision(obj1, obj3) {
		t.Errorf("Failed test collision")
	}
}
