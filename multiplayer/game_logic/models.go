package game_logic

// Эй, детка, ты модель или подделка?
type Game struct {
	GameObjects GameObjects `json:"gameobjects"`
	Map         Map         `json:"map"`
	Wave        int         `json:"wave"`
	Url         string      `json:"url"`
	Stage       string      `json:"stage"`
}

type InputMessage struct {
	Up      bool    `json:"up"`
	Down    bool    `json:"down"`
	Left    bool    `json:"reft"`
	Right   bool    `json:"right"`
	Angular float32 `json:"ang"`
	Shot    bool    `json:"shot"`
}

type DynamycObject struct {
	Name string

	Hp    float32 `json:"hp"`
	HpCap float32 `json:"hpcap"`

	X     float32 `json:"x"`
	Y     float32 `json:"y"`
	Xsize float32 `json:"xsize"`
	Ysize float32 `json:"ysize"`

	Velocity       float32 `json:"velocity"`
	VelocityBarior float32
}

type Bullet struct {
	Object   *DynamycObject `json:"object"`
	Damage   float32
	PlayerId int8 `json:"playerid"`
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
	Angular    int            `json:"ang"`
	Weapon     *Weapon
	//Inventory 		[]DynamycObject
}

type Adv struct {
	Object  *DynamycObject `json:"object"`
	Url     string
	Pict    string `json:"pict"`
	XTarget float32
	YTarget float32
	Angular float32 `json:"ang"`
}

type Barior struct {
	Id     int            `json:"id"`
	Object *DynamycObject `json:"object"`
}

type Map struct {
	TileSize int   `json:"tailsize"`
	SizeX    int   `json:"sizex"`
	SizeY    int   `json:"sizey"`
	Field    []int `json:"field"`
}

type GameObjects struct {
	Players []Player `json:"palyers"`
	Advs    []Adv    `json:"advs"`
	Bullets []Bullet `json:"bullets"`
	Bariors []Barior `json:"bariors"`
}
