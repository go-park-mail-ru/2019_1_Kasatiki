package game_logic

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

func SimpleCollisionEvent(obj1, obj2 *DynamycObject, moves Moves) {
	if moves.Up || moves.Down {
		if obj1.Y-(obj2.Y-obj2.Ysize) > obj2.Y+obj2.Ysize-obj1.Y && (obj1.X < obj2.X-obj1.Xsize && obj1.X > obj2.X+obj2.Xsize) {
			obj1.Y = obj2.Y + obj2.Ysize
		} else {
			obj1.Y = obj2.Y - obj1.Ysize
		}
	}
	if moves.Left || moves.Right {
		if obj1.X-(obj2.X-obj2.Xsize) > obj2.X+obj2.Xsize-obj1.X {
			obj1.X = obj2.X + obj2.Xsize
		} else {
			obj1.X = obj2.X - obj1.Xsize
		}
	}
}

func (p1 *Player) PlayerToPlayer(p2 *DynamycObject, moves Moves) {
	SimpleCollisionEvent(p1.Object, p2, moves)
}

func IsCollision(obj1, obj2 *DynamycObject) bool {
	if obj1.Y > obj2.Y-obj1.Ysize && // граница	сверху
		obj1.Y < obj2.Y+obj2.Ysize && // граница	снизу
		obj1.X > obj2.X-obj1.Xsize && // граница	справа
		obj1.X < obj2.X+obj2.Xsize { // граница 	слева
		//fmt.Println("Colision with Obj_1: ", obj1.Name, " and  Obj_2: ", obj2.Name)
		return true
	}
	return false
}

func (g *Game) CollectObjectsForPlayer(nickname string) []*DynamycObject {
	var objs []*DynamycObject
	for k, v := range g.GameObjects.Players {
		if k != nickname {
			objs = append(objs, v.Object)
		}
	}
	for _, v := range g.GameObjects.Barrier {
		objs = append(objs, v.Object)
	}
	return objs
}

// Входная точка для изменения состояния игры
// Принимает в себя структуру, которая получилась после разкодирования из json
func (g *Game) EventListener(mes InputMessage, nickname string) (res GameStatus) {
	g.GameObjects.Players[nickname].SetAngular(mes.Angular)
	delta := g.GameObjects.Players[nickname].Object.Velocity
	//var moves Moves
	objs := g.CollectObjectsForPlayer(nickname)
	if mes.Down {
		g.GameObjects.Players[nickname].Object.Y += delta
		for _, obj := range objs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				g.GameObjects.Players[nickname].Object.Y -= delta
			}
		}
	}

	if mes.Up {
		g.GameObjects.Players[nickname].Object.Y -= delta
		for _, obj := range objs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				g.GameObjects.Players[nickname].Object.Y += delta
			}
		}
	}
	if mes.Left {
		g.GameObjects.Players[nickname].Object.X -= delta
		for _, obj := range objs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				g.GameObjects.Players[nickname].Object.X += delta
			}
		}

		//moves.Left = true
	}
	if mes.Right {
		g.GameObjects.Players[nickname].Object.X += delta
		for _, obj := range objs {
			// Если произошла коллизия
			if IsCollision(g.GameObjects.Players[nickname].Object, obj) {
				g.GameObjects.Players[nickname].Object.X -= delta
			}
		}
		//moves.Right = true
	}
	if mes.Shot {
		g.GameObjects.Players[nickname].Shot()
	}

	// Здесь обрабатываем коллизии игрока, который только сходил:

	// 1) Собираем все объекты для этих коллизий (другие игроки, барьеры

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
	for _, adv := range g.GameObjects.Advs {
		adv.MoveToPlayer(g.Map)
		var info AdvInfo
		info.Object = adv.Object
		res.Advs = append(res.Advs, info)
	}
	return
}
