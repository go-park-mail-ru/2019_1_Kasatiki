package game_logic

import (
	"math"
)

func (b *Bullet) Run() {
	b.Object.X += int(float64(b.Object.Velocity-3) * math.Cos(float64(b.Angle)) * 0.5)
	b.Object.Y += int(float64(b.Object.Velocity-3) * math.Sin(float64(b.Angle)) * 0.5)
}
