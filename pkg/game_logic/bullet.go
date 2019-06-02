package game_logic

import (
	"log"
	"math"
)

func (b *Bullet) IsOnMap(m *Map) bool {
	if b.Object.X < 0 || b.Object.Y < 0 ||
		b.Object.X > m.TileSize*m.SizeX ||
		b.Object.Y > m.TileSize*m.SizeY {
		return false
	}

	return true
}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func (bullet *Bullet) IsCollisionInWay(obj *DynamycObject) bool {
	// x, y - его будущие координаты после перемещения (за один тик)
	newX := float64(bullet.Object.X + int(float64(bullet.Object.Velocity-3)*math.Cos(float64(bullet.Angle))*0.5))
	newY := float64(bullet.Object.Y + int(float64(bullet.Object.Velocity-3)*math.Sin(float64(bullet.Angle))*0.5))
	// s-Расстояние до новой координаты
	s := distance(float64(bullet.Object.X), float64(bullet.Object.Y), newX, newY)
	// Помещаем координаты объекта в центр теперь обхект-> окружность
	circleX := float64(obj.X + obj.Xsize/2)
	circleY := float64(obj.Y + obj.Ysize/2)
	// будем считать рекламу окружностью радиусом size / 2 < R < size * sqrt(2)
	objRadius := float64(obj.Xsize/2) * 1.3
	// a-Расстояние от центра окружности до предыдущей точки
	a := distance(float64(obj.X), float64(obj.Y), circleX, circleY)
	// b=Расстояние от центра окружности до точки перемещения пули (за один тик)
	b := distance(newX, newY, circleX, circleY)
	// h-расстояние от центра до прямой, содержащей prev и next координаты пули
	// Но сначало нужно проверить выражение под корнем
	h := a*a - ((b*b-s*s-a*a)/(2*s))*((b*b-s*s-a*a)/(2*s))
	if h < 0 {
		return false
	} else {
		h = math.Sqrt(h)
	}
	// Если расстояние до прямой меньше радиуса, то коллизия
	if h < objRadius {
		return true
	}
	log.Println(h)
	return false
}

func (b *Bullet) Run() {
	b.Object.X += int(float64(b.Object.Velocity-3) * math.Cos(float64(b.Angle)) * 0.5)
	b.Object.Y += int(float64(b.Object.Velocity-3) * math.Sin(float64(b.Angle)) * 0.5)
}
