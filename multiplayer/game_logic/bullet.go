package game_logic

import (
	"math"
)

func (b *Bullet) Go(a float64) {
	for {
		b.Object.X += b.Object.Velocity * (int)(math.Cos(a))
		b.Object.Y += b.Object.Velocity * (int)(math.Sin(a))
	}
}