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

func (adv *Adv) MoveToPlayer(m *Map) {
	player := adv.Player
	// Тангенс угла наклона
	angular := math.Atan2(float64(player.Object.Y-adv.Object.Y), float64(player.Object.X-adv.Object.X))
	distanceX := int(float64(adv.Object.Velocity) * math.Cos(angular))
	distanceY := int(float64(adv.Object.Velocity) * math.Sin(angular))
	if angular > 0.0 {
		if angular <= math.Pi/2 {
			if m.Field[(adv.Object.Y+distanceY)/m.TileSize+1][adv.Object.X/m.TileSize] == 1 {
				distanceX = adv.Object.Velocity
				distanceY = 0
				if m.Field[adv.Object.Y/m.TileSize][(adv.Object.X+distanceX)/m.TileSize+1] == 1 {
					distanceX = 0
				}
			} else {
				if m.Field[adv.Object.Y/m.TileSize][(adv.Object.X+distanceX)/m.TileSize+1] == 1 {
					distanceY = adv.Object.Velocity
					distanceX = 0
				}
			}
		} else {
			if m.Field[(adv.Object.Y+distanceY)/m.TileSize+1][(adv.Object.X+distanceX)/m.TileSize+1] == 1 {
				distanceX = -adv.Object.Velocity
				distanceY = 0
				if m.Field[adv.Object.Y/m.TileSize][(adv.Object.X+distanceX)/m.TileSize-1] == 1 {
					distanceX = 0
				}
			} else {
				if m.Field[adv.Object.Y/m.TileSize][(adv.Object.X+distanceX)/m.TileSize] == 1 {
					distanceY = adv.Object.Velocity
					distanceX = 0
				}
			}
		}
	} else {
		// log.Println(distanceY, distanceX)
		// log.Println("HELLO", time.Now())
		if angular >= -math.Pi/2 {
			if m.Field[(adv.Object.Y+distanceY)/m.TileSize-1][adv.Object.X/m.TileSize] == 1 {
				distanceX = adv.Object.Velocity
				distanceY = 0
				if m.Field[(adv.Object.Y+distanceY)/m.TileSize+1][(adv.Object.X+distanceX)/m.TileSize+1] == 1 {
					distanceX = 0
				}
			} else {
				if m.Field[adv.Object.Y/m.TileSize+1][(adv.Object.X+distanceX)/m.TileSize+1] == 1 {
					distanceY = -adv.Object.Velocity
					distanceX = 0
				}
			}
		} else {
			if m.Field[(adv.Object.Y+distanceY)/m.TileSize][(adv.Object.X+distanceX)/m.TileSize+1] == 1 {
				distanceX = -adv.Object.Velocity
				distanceY = 0
				if m.Field[adv.Object.Y/m.TileSize+1][(adv.Object.X+distanceX)/m.TileSize-1] == 1 {
					distanceX = 0
				}
			} else {
				if m.Field[adv.Object.Y/m.TileSize][(adv.Object.X+distanceX)/m.TileSize] == 1 {
					distanceY = -adv.Object.Velocity
					distanceX = 0
				}
			}
		}
	}
	adv.Object.Y += distanceY
	adv.Object.X += distanceX
	// if adv.Object.X != player.Object.X {
	// 	adv.Object.X += int(float64(adv.Object.Velocity) * math.Cos(angular))
	// }
	// if adv.Object.Y != player.Object.Y {
	// 	adv.Object.Y += int(float64(adv.Object.Velocity) * math.Sin(angular))
	// }
	// log.Println("P", player.Object.Y, player.Object.X)
	// log.Println("A", adv.Object.Y, adv.Object.X)
}
