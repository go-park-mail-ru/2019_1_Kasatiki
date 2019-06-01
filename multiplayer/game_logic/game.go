package game_logic

// import "fmt"

// Колизии

// Логика такая - сущность переместилась, проверяем для нее кализии
// Игрок - барьер
// Игрок - Игрок
// Реклама - Игрок
// Реклама - блок
// Пуля - Реклама
// Пуля - Блок
// Пуля - Игрок

func (g *Game) Update() {
	// Передвигаем рекламу
	// Проверяем колизии рекламы
	// Передвигаем пули
	// Проверяем колизии пули

	//g.CollisionMeneger()

}

// Проверяет, уничтожены ли все объекты
// Если да - заканчивает волну и игра переходит в паузу
func (g *Game) IsEndOfWave() {

}

// Todo Сделать Динамическую мапу объектов, в которой будут лежать игроки, рекламы и пули после (актуально для рекламы и пуль) но это не точно
func (g *Game) CollectObjectsForPlayer(nickname string, player *DynamycObject) []*DynamycObject {
	var objs []*DynamycObject

	for k, v := range g.GameObjects.Players {
		if k != nickname {
			objs = append(objs, v.Object)
		}
	}
	var count int
	numbs := []int{}
	for _, z := range g.Zones {
		// Создаем DynamicObject из зоны
		zoneObj := &DynamycObject{
			// Name:  fmt.Sprintf("Zone %d", i),
			X:     z.StartX,
			Y:     z.StartY,
			Xsize: z.EndX - z.StartX,
			Ysize: z.EndY - z.StartY,
		}
		// Если игрок в зоне, то заносим объекты из зоны в слайс вероятных колизий
		if IsCollision(player, zoneObj) {
			count++
			numbs = append(numbs, z.Number)
			objs = append(objs, g.StaticCollection[z.Number]...)
		}
	}

	//fmt.Println("--------------------------------------------------------------------")
	//for _,n  := range numbs {
	//	fmt.Println("colision zone is ", n)
	//}
	//fmt.Println("--------------------------------------------------------------------")
	//fmt.Println("Число Зон вошедших с игроком в коллизию",count)

	//fmt.Println("Before objs:", len(objs))
	//objs = ObjectFilter(objs)
	//
	//fmt.Println("--------------------------------------------------------------------")
	//fmt.Println("After Objs: ", len(objs))
	//for _, b := range objs {
	////
	//	fmt.Printf("Name : %s, StartX : %d, StartY : %d, EndX : %d, EndY : %d \n", b.Name, b.X, b.Y, b.Xsize, b.Ysize)
	//}
	//fmt.Println("--------------------------------------------------------------------")
	return objs
}

var gameObjs []*DynamycObject

func GetGameObjs() []*DynamycObject {
	return gameObjs
}

// Входная точка для изменения состояния игры
// Принимает в себя структуру, которая получилась после разкодирования из json
func (g *Game) EventListener(mes InputMessage, nickname string) (res GameStatus) {

	g.GameObjects.Players[nickname].SetAngular(mes.Angular)
	delta := g.GameObjects.Players[nickname].Object.Velocity + 5

	// Собирем объекты с которыми игрок вероятнее всего столнется
	gameObjs = g.CollectObjectsForPlayer(nickname, g.GameObjects.Players[nickname].Object)

	if mes.Down {
		g.GameObjects.Players[nickname].Object.Y += delta
		for _, obj := range gameObjs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {

				// fmt.Println("Colision with Obj_1: ", g.GameObjects.Players[nickname].Object.Name, " and  Obj_2: ", obj.Name)
				g.GameObjects.Players[nickname].Object.Y -= delta
			}
		}
	}

	if mes.Up {
		g.GameObjects.Players[nickname].Object.Y -= delta
		for _, obj := range gameObjs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				// fmt.Println("Colision with Obj_1: ", g.GameObjects.Players[nickname].Object.Name, " and  Obj_2: ", obj.Name)
				g.GameObjects.Players[nickname].Object.Y += delta
			}
		}
	}
	if mes.Left {
		g.GameObjects.Players[nickname].Object.X -= delta
		for _, obj := range gameObjs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				// fmt.Println("Colision with Obj_1: ", g.GameObjects.Players[nickname].Object.Name, " and  Obj_2: ", obj.Name)
				// fmt.Printf("Obj_1 : X = %d , Y = %d , Xsize = %d, Ysize = %d \n", g.GameObjects.Players[nickname].Object.X, g.GameObjects.Players[nickname].Object.Y, g.GameObjects.Players[nickname].Object.Xsize, g.GameObjects.Players[nickname].Object.Ysize)
				// fmt.Printf("Obj_2 : X = %d , Y = %d , Xsize = %d, Ysize = %d \n", obj.X, obj.Y, obj.Xsize, obj.Ysize)
				g.GameObjects.Players[nickname].Object.X += delta
			}
		}

		//moves.Left = true
	}
	if mes.Right {
		g.GameObjects.Players[nickname].Object.X += delta
		for _, obj := range gameObjs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				// fmt.Println("Colision with Obj_1: ", g.GameObjects.Players[nickname].Object.Name, " and  Obj_2: ", obj.Name)
				g.GameObjects.Players[nickname].Object.X -= delta
			}
		}
		//moves.Right = true
	}
	if mes.Shot {
		bullet := g.GameObjects.Players[nickname].Shot(mes.Angular)
		g.GameObjects.Bullets = append(g.GameObjects.Bullets, bullet)
	}

	// Здесь обрабатываем коллизии игрока, который только сходил:

	//
	// d1) Собираем все объекты для этих коллизий (другие игроки, барьеры

	// 2) Проверяем этого игрока на колизии с этими объектами
	//for _, obj := range objs {
	//		// Если произошла коллизия
	//		if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
	//			g.GameObjects.Players[nickname].PlayerToPlayer(obj, moves)
	//		}
	//}

	for _, p := range g.GameObjects.Players {
		var info PlayerInfo
		info.Object = p.Object
		info.Id = p.Id
		info.CashPoints = p.CashPoints
		info.Nickname = p.Nickname
		info.Id = p.Id
		res.Players = append(res.Players, info)
	}

	// Reklama
	for _, adv := range g.GameObjects.Advs {
		// log.Println(adv.Object.X, adv.Object.Y)
		adv.MoveToPlayer(g.Map)
		var info AdvInfo
		info.Object = adv.Object
		res.Advs = append(res.Advs, info)
	}

	return
}
