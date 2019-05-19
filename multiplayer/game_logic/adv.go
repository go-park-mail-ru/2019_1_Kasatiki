package game_logic

import (
	"math"
)

func (adv *Adv) Spawn(x int, y int) {
	adv.Object = &DynamycObject{
		Name:     "Adv",
		Hp:       10,
		X:        x,
		Y:        y,
		Velocity: 2,
	}
}

func (adv *Adv) MoveToPlayer() {
	player := adv.Player
	// Тангенс угла наклона
	angular := math.Atan2(float64(player.Object.Y - adv.Object.Y), float64(player.Object.X - adv.Object.X))
	if adv.Object.X != player.Object.X {
		adv.Object.X += int(float64(adv.Object.Velocity) * math.Cos(angular))
	}
	if adv.Object.Y != player.Object.Y {
		adv.Object.Y += int(float64(adv.Object.Velocity) * math.Sin(angular))
	}
	// if adv.Object.X > player.Object.X {
	// 	adv.Object.X -= int(math.Abs(float64(adv.Object.Velocity) * math.Cos(angular)))
	// } else if adv.Object.X < player.Object.X {
	// 	adv.Object.X += int(math.Abs(float64(adv.Object.Velocity) * math.Cos(angular)))
	// }
	// if adv.Object.Y > player.Object.Y {
	// 	adv.Object.Y -= int(math.Abs(float64(adv.Object.Velocity) * math.Sin(angular)))
	// } else if adv.Object.Y < player.Object.Y {
	// 	log.Println("  EHUU")
	// 	adv.Object.Y += int(math.Abs(float64(adv.Object.Velocity) * math.Sin(angular)))
	// }
}