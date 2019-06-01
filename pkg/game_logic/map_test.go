package game_logic

import "testing"

func TestMapGeneration(t *testing.T) {
	m, _ := MapGeneration()
	if m == nil {
		t.Errorf("Failed map generation")
	}
}

func TestInsert(t *testing.T) {
	var s []int
	s = append(s, 1, 2, 3)
	s = Insert(s, 1, 9)
	if s[1] != 9 {
		t.Errorf("Failed insert test")
	}
}
