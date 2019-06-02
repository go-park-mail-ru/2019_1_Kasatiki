package game_logic

import (
	"log"
	"math"
)

func (adv *Adv) Spawn(x int, y int, size int) {
	adv.Object = &DynamycObject{
		Name:     "Adv",
		Hp:       10,
		HpCap:    10,
		X:        x,
		Y:        y,
		Velocity: 2,
		Xsize:    size,
		Ysize:    size,
	}
}

func (adv *Adv) MoveWithWay(way Points, m *Map) {
	if len(way) <= 1 { // way[0] - коордната самой рекламы
		return
	}
	distance := adv.Object.Velocity
	// Количество совершенных шагов по карте (100x100)
	//  way - массив, включающий координаты от рекламы к цели(игроку);
	// len(way) - 1 = количество шагов до цеели
	approximateStepCount := distance / m.TileSize
	// Иначе двигаем рекламу на approximateStepCount по мапе
	// + остается кусочек, на который еще нужно додвинуть рекламу
	distanceToNearestCell := int(math.Sqrt(
		float64((way[1].XCell*m.TileSize-adv.Object.X)*(way[1].XCell*m.TileSize-adv.Object.X) +
			(way[1].YCell*m.TileSize-adv.Object.Y)*(way[1].YCell*m.TileSize-adv.Object.Y))))
	// if approximateStepCount > 0 || distanceToNearestCell < distance {
	// 	distance -= approximateStepCount*m.TileSize - distanceToNearestCell
	// }
	if approximateStepCount > 0 {
		distance -= approximateStepCount * m.TileSize
	}
	if distance >= m.TileSize {
		approximateStepCount += 1
		distance -= m.TileSize
	}
	if approximateStepCount >= len(way)-1 { // Если реклама достает до игрока за 1 шаг
		adv.Object.X = adv.Player.Object.X
		adv.Object.Y = adv.Player.Object.Y
		return
	}
	// столько он пройдет до поля way[appr..].xcell way[appr..].ycell
	log.Println(approximateStepCount, "---hello", distance, distanceToNearestCell, len(way), adv.Object.Velocity)
	log.Println(adv.Object.X, adv.Object.Y, adv.Player.Object.X, adv.Player.Object.Y)

	if approximateStepCount > 0 {
		adv.Object.X = way[approximateStepCount].XCell * m.TileSize
		adv.Object.Y = way[approximateStepCount].YCell * m.TileSize
	}
	if way[approximateStepCount].XCell == way[approximateStepCount+1].XCell {
		if way[approximateStepCount].YCell > way[approximateStepCount+1].YCell {
			adv.Object.Y -= distance
		} else if way[approximateStepCount].YCell < way[approximateStepCount+1].YCell {
			adv.Object.Y += distance
		}
	} else {
		if way[approximateStepCount].XCell > way[approximateStepCount+1].XCell {
			adv.Object.X -= distance
		} else {
			adv.Object.X += distance
		}
	}
}

func (adv *Adv) MoveWithWay_with_one_step(way Points, m *Map) {
	if len(way) <= 1 {
		return
	}
	distance := adv.Object.Velocity
	distanceToNearestCell := int(math.Sqrt(
		float64((way[len(way)-2].XCell*m.TileSize-adv.Object.X)*(way[len(way)-2].XCell*m.TileSize-adv.Object.X) +
			(way[len(way)-2].YCell*m.TileSize-adv.Object.Y)*(way[len(way)-2].YCell*m.TileSize-adv.Object.Y))))
	if distanceToNearestCell <= distance {
		adv.Object.X = way[len(way)-2].XCell * m.TileSize
		adv.Object.Y = way[len(way)-2].YCell * m.TileSize
		return
	}
	if way[len(way)-1].XCell == way[len(way)-2].XCell {
		if way[len(way)-1].YCell > way[len(way)-2].YCell {
			adv.Object.Y -= distance
		} else if way[len(way)-1].YCell < way[len(way)-2].YCell {
			adv.Object.Y += distance
		}
	} else {
		if way[len(way)-1].XCell > way[len(way)-2].XCell {
			adv.Object.X -= distance
		} else if way[len(way)-1].XCell < way[len(way)-2].XCell {
			adv.Object.X += distance
		}
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
}
