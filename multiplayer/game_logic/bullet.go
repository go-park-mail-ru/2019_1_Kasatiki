package game_logic

import (
	"math"
)

func (b *Bullet) Run() {
	b.Object.X += b.Object.Velocity * int(math.Cos(float64(b.Angle)))
	b.Object.Y += b.Object.Velocity * int(math.Sin(float64(b.Angle)))
}