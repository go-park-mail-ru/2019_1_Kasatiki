package game_logic

//	Собирает все объекты для проверки на колизии
func (g *Game) DynamicObjectsCollector() (objs []DynamycObject) {
	return
}

// Управляет исходом колизий
func CollisionMeneger(objects []DynamycObject) {

}

// Проверяет, уничтожены ли все объекты
// Если да - заканчивает волну и игра переходит в паузу
func (g *Game) IsEndOfWave() {

}

// Входная точка для изменения состояния игры
// Принимает в себя структуру, которая получилась после разкодирования из json
func (g *Game) EventListener(mes InputMessage, nickname string) (res GameStatus) {
	g.GameObjects.Players[nickname].SetAngular(mes.Angular)
	delta := g.GameObjects.Players[nickname].Object.Velocity
	if mes.Down {
		g.GameObjects.Players[nickname].Object.Y += delta
	}
	if mes.Up {
		g.GameObjects.Players[nickname].Object.Y -= delta
	}
	if mes.Left {

		g.GameObjects.Players[nickname].Object.X -= delta
	}
	if mes.Right {
		g.GameObjects.Players[nickname].Object.X += delta
	}

	if mes.Shot {
		g.GameObjects.Players[nickname].Shot()
	}

	for _, p := range g.GameObjects.Players {
		var info PlayerInfo
		info.Object = p.Object
		info.Id = p.Id
		info.CashPoints = p.CashPoints
		info.Nickname = p.Nickname
		info.Id = p.Id
		res.Players = append(res.Players, info)
	}
	return
}
