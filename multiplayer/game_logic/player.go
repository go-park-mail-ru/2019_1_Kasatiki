package game_logic

import (
// "fmt"
)

// Хардкод дефолного оружия
func (w *Weapon) SetBullet(dam float32, player Player) {
	w.Bullet = &Bullet{}
	w.Bullet.Object = &DynamycObject{
		Name:     w.Name,
		X:        player.Object.X,
		Y:        player.Object.Y,
		Velocity: 7,
	}
	w.Bullet.Damage = dam
	w.Bullet.PlayerId = player.Id
}

func (p *Player) CreateDefaultWeapon() {
	p.Weapon = &Weapon{
		Id:       0,
		Name:     "Deagle",
		FireRate: 3,
		Coast:    300,
	}
	p.Weapon.SetBullet(20, *p)
}

//// Обращение к бд для смены оружия(покупка)
//func (p *Player) ChangeWeapon() (w *Weapon) {
//	return
//}

func (p *Player) Spawn(x int, y int, sizeX int, sizeY int) {
	p.Object = &DynamycObject{
		Name:     "Player",
		Hp:       100,
		X:        x,
		Y:        y,
		Velocity: 5,
		Xsize:    sizeX,
		Ysize:    sizeY,
	}
}

//// Создание пули
//func (p *Player) BulletsCreate() (bs []Bullet) {
//	return
//}

// аппендит пули в массив пулей
func (p *Player) Shot(a float32) *Bullet {

	var b Bullet

	b.Object = &DynamycObject{
		Name:     "Bullet",
		Hp:       0,
		X:        p.Object.X + p.Object.Xsize/2,
		Y:        p.Object.Y + p.Object.Ysize/2,
		Velocity: 15,
		Xsize:    5,
		Ysize:    5,
	}

	b.Damage = 5
	b.Angle = a

	return &b
}

func (p1 *Player) PlayerToPlayer(p2 *DynamycObject, moves Moves) {
	SimpleCollisionEvent(p1.Object, p2, moves)
}

func (p *Player) SetAngular(ang float32) {
	p.Angular = ang
}

func (p *Player) SetNickname(nickname string) {
	p.Nickname = nickname
}
