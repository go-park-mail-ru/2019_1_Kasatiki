package game_logic



//	Пример разбияния карты на зоны
//	при ZoneOnXAxis = 4 и ZoneOnYAxis = 4
//
//
//	╔════╦════╦════╦════╗
//	║    ║    ║    ║    ║
//	║ 1  ║ 5  ║ 9  ║ 13 ║
//	╠════╬════╬════╬════╣
//	║    ║    ║    ║    ║
//	║ 2  ║ 6  ║ 10 ║ 14 ║
//	╠════╬════╬════╬════╣
//	║    ║    ║    ║    ║
//	║ 3  ║ 7  ║ 11 ║ 15 ║
//	╠════╬════╬════╬════╣
//	║    ║    ║    ║    ║
//	║ 4  ║ 8  ║ 12 ║ 16 ║
//	╚════╩════╩════╩════╝



// Создаем зоны и заносим все блоки
func (g *Game) ZonesIni() {
	// Число разбиений по оси X
	const ZoneOnXAxis = 6
	// Число разбиений по си Y
	const ZoneOnYAxis = 6

	// Размер шага по соответствующим осям
	deltaX := g.Map.SizeX * g.Map.TileSize / ZoneOnXAxis
	deltaY := g.Map.SizeY * g.Map.TileSize / ZoneOnYAxis

	// Число зон, в которых будем искать объекты
	ZoneNumbers := ZoneOnXAxis * ZoneOnYAxis

	startX := 0
	startY := 0
	counter := 0
	for i := 0; i < ZoneNumbers + ZoneOnXAxis; i++ {

		var z Zone
		z.StartX = startX
		z.StartY = startY
		z.EndY =  startY + deltaY
		z.EndX = startX + deltaX

		// Если еще не достигли конца карты по Y
		if z.EndY <= g.Map.SizeY * g.Map.TileSize {
			// Наращиваем счетчик
			counter ++
			// Опускаемся на одну клетку вниз(для следующей иттерации)
			startY = z.EndY
			z.Number = counter
			// Отправляем зону в слайс
			g.Zones = append(g.Zones, &z)
			continue
		} else { // если достигли конца
			// Смещаемся вправо на одну клетку(для следующей иттерации)
			startX = z.EndX
			// Смещаемся на самый верх карты
			startY = 0
			continue
		}
	}

	// Мапа в которой в качестве ключа используется номер зоны
	// В качестве значения - массив динамических объектов
	zones := make(map[int][]*DynamycObject, ZoneNumbers)

	// Пробегаемся по всем зонам и проверяем входят ли в них барьеры
	for _, z := range g.Zones {
		var objs []*DynamycObject

		// Создаем DynamicObject из зоны
		zoneObj := &DynamycObject {
			Name: "Zone",
			X : z.StartX,
			Y : z.StartY,
			Xsize : z.EndX - z.StartX,
			Ysize : z.EndY - z.StartY,
		}
		// Проходимся по всем барьерам, если он задевает зону
		// Заносим в слайс соответсвующей зоны
		for _, v := range g.GameObjects.Barrier {
			if IsCollision(zoneObj, v.Object) {
				objs = append(objs, v.Object)
			}
		}
		zones[z.Number] = objs
	}
	g.StaticCollection = zones
	return
}
