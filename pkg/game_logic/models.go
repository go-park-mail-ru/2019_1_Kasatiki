package game_logic

// Эй, детка, ты модель или подделка?
type Game struct {
	GameObjects      *GameObjects `json:"gameobjects"`
	Map              *Map         `json:"map"`
	Wave             int          `json:"wave"`
	Url              string       `json:"url"`
	Stage            string       `json:"stage"`
	Zones            []*Zone
	StaticCollection map[int][]*DynamycObject
}

type Zone struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
	Number int
}

type BulletStatus struct {
	Bullets []*Bullet `json:"bullets"`
}

type GameStatus struct {
	Players []PlayerInfo `json:"players"`
	Advs    []AdvInfo    `json:"advs"`
}

type StartGame struct {
	Map     Map          `json:"map"`
	Barrier []*Barrier   `json:"barriers"`
	Players []PlayerInfo `json:"players"`
	Advs    []AdvInfo    `json:"advs"`
	Id      int          `json:"id"`
}

type PlayerInfo struct {
	Object     *DynamycObject `json:"object"`
	CashPoints float32        `json:"cash"`
	Nickname   string         `json:"nickname"`
	Id         int            `json:"id"`
	Angular    float32        `json:"ang"`
}

type AdvInfo struct {
	Object *DynamycObject `json:"object"`
}

type InputMessage struct {
	Up      bool    `json:"up"`
	Down    bool    `json:"down"`
	Left    bool    `json:"left"`
	Right   bool    `json:"right"`
	Angular float32 `json:"angle"`
	Shot    bool    `json:"shot"`
}

type DynamycObject struct {
	Name string

	Hp    float32 `json:"hp"`
	HpCap float32 `json:"hpcap"`

	X     int `json:"x"`
	Y     int `json:"y"`
	Xsize int `json:"xsize"`
	Ysize int `json:"ysize"`

	Velocity       int `json:"velocity"`
	VelocityBarior int
}

type Bullet struct {
	Object *DynamycObject `json:"object"`
	Damage float32
	Angle  float32
	Player *Player `json:"-"`
}

type Weapon struct {
	Id       int
	Name     string
	FireRate float32
	Magazine int
	Bullet   *Bullet
	Coast    int
}

//type Item struct {}

type Player struct {
	Object     *DynamycObject `json:"object"`
	CashPoints float32        `json:"cash"`
	Nickname   string         `json:"nickname"`
	Id         int            `json:"id"`
	Angular    float32        `json:"ang"`
	Weapon     *Weapon
	Stepback   *Moves
	Killed     int `json:"killed"`
	//Inventory 		[]DynamycObject
}

// Нужна для отката движений игрока
type Moves struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type Adv struct {
	Object *DynamycObject `json:"object"`
	Id     int
	Url    string
	Pict   string `json:"pict"`
	Name   string
	//XTarget float32
	//YTarget float32
	//Angular float32 `json:"ang"`
	Player *Player
}

type Barrier struct {
	Id     int            `json:"id"`
	Object *DynamycObject `json:"object"`
}

type Map struct {
	TileSize int           `json:"tailsize"`
	SizeX    int           `json:"sizex"`
	SizeY    int           `json:"sizey"`
	Field    [100][100]int `json:"field"`
}

type GameObjects struct {
	Players map[string]*Player `json:"palyers"`
	Advs    []*Adv             `json:"advs"`
	Bullets []*Bullet          `json:"bullets"`
	Barrier []*Barrier         `json:"barrier"`
}

type Point struct {
	XCell int
	YCell int

	G int // g(x). Стоимость пути от начальной вершины. У start g(x) = 0
	F int // f(x) = g(x) + h(x)
}

type Points []*Point
