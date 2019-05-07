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
func (g *Game) EventListener(mes InputMessage, nickname string) {

}
