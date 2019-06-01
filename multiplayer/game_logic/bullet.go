package game_logic

import (
	"math"
	// "fmt"
)

func (b *Bullet) IsOnMap(m *Map) bool {
	if b.Object.X < 0 || b.Object.Y < 0 ||
		b.Object.X > m.TileSize*m.SizeX ||
		b.Object.Y > m.TileSize*m.SizeY {
		return false
	}

	return true
}

func (b *Bullet) Run() {
	b.Object.X += int(float64(b.Object.Velocity-3) * math.Cos(float64(b.Angle)) * 0.5)
	b.Object.Y += int(float64(b.Object.Velocity-3) * math.Sin(float64(b.Angle)) * 0.5)
	// fmt.Println(int(math.Cos(float64(b.Angle))), math.Cos(float64(b.Angle)), b.Object.X, b.Object.Y)
}
