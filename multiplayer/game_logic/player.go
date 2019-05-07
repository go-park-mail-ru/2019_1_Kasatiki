package game_logic

// Хардкод дефолного оружия
func (p *Player) CreateDefaultWeapon() (w *Weapon) {
	return
}

// Обращение к бд для смены оружия(покупка)
func (p *Player) ChangeWeapon() (w *Weapon) {
	return
}

// Меняет состоние координат внутри поля Объект
func (p *Player) Move() {
	//wasd
}

// Создание пули
func (p *Player) BulletsCreate() (bs []Bullet) {
	return
}

// аппендит пули в массив пулей
func (p *Player) Shot() {

}
