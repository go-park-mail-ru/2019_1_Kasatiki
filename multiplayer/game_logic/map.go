package game_logic

import (
	"math/rand"
	"time"
)

type Template struct {
	Tmp      []int
	Barriers [][]int
}

func Insert(slice []int, index, value int) []int {
	// Увеличиваем срез на один элемент
	slice = slice[0 : len(slice)+1]
	// Используем copy для перемещения верхней части среза наружу и создания пустого места
	copy(slice[index+1:], slice[index:])
	// Записываем новое значение.
	slice[index] = value
	// Возвращаем результат.
	return slice
}

// Создание карты
func MapGeneration() (*Map, []*Barrier) {

	var m Map
	var b []*Barrier

	// Инициализируем параметры карты
	m.TileSize = 10
	m.SizeX = 100
	m.SizeY = 100

	// Логика заполнения карты препятствиями:
	// Делим карту на 16 блоков (4x4 каждый по 25 тайлов)
	// Заполняем каждый блок препятствиями

	// blockX := ( m.SizeX  ) / blockCount;
	// blockY := ( m.SizeY ) / blockCount;

	// Создаем границы карты

	// Генерируем вертикальные границы
	// for i := 0; i < m.SizeY; i++ {
	//  for j := 0; j < m.SizeX; j++ {
	//      if i == 0 || i == m.SizeY-1 || j == 0 || j == m.SizeX -1 {
	//             m.Field[m.SizeY * i + j] = 1;
	//         }
	//  }
	// }

	// Задаем массив шаблонов карт:
	// Каждый шаблон - массив 20x20, заполенный препядствием

	tmp1 := Template{
		Tmp: []int{
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		},
		Barriers: [][]int{
			[]int{
				7 * m.TileSize, 0, 3 * m.TileSize, 2 * m.TileSize,
			},
			[]int{
				0, 4 * m.TileSize, 6 * m.TileSize, 3 * m.TileSize,
			},
			[]int{
				0, 4 * m.TileSize, 6 * m.TileSize, 3 * m.TileSize,
			},
		},
	}

	tmp2 := Template{
		Tmp: []int{
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 1, 1, 1, 0, 0, 0, 0,
		},
		Barriers: [][]int{
			[]int{
				0, 0, 6 * m.TileSize, 3 * m.TileSize,
			},
			[]int{
				3 * m.TileSize, 7 * m.TileSize, 3 * m.TileSize, 3 * m.TileSize,
			},
		},
	}

	tmp3 := Template{
		Tmp: []int{
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		Barriers: [][]int{
			[]int{
				0, 0, 3 * m.TileSize, 2 * m.TileSize,
			},
			[]int{
				0, 4 * m.TileSize, 6 * m.TileSize, 3 * m.TileSize,
			},
		},
	}

	tmp4 := Template{
		Tmp: []int{
			0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
			0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
			0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		},
		Barriers: [][]int{
			[]int{
				4 * m.TileSize, 0, 6 * m.TileSize, 3 * m.TileSize,
			},
			[]int{
				7 * m.TileSize, 3 * m.TileSize, 3 * m.TileSize, 3 * m.TileSize,
			},
			[]int{
				0, 7 * m.TileSize, 3 * m.TileSize, 3 * m.TileSize,
			},
		},
	}

	tmp5 := Template{
		Tmp: []int{
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
			0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
			0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
			0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
		},
		Barriers: [][]int{
			[]int{
				0, 0, 2 * m.TileSize, 3 * m.TileSize,
			},
			[]int{
				4 * m.TileSize, 4 * m.TileSize, 6 * m.TileSize, 3 * m.TileSize,
			},
			[]int{
				7 * m.TileSize, 7 * m.TileSize, 3 * m.TileSize, 3 * m.TileSize,
			},
		},
	}

	tmp6 := Template{
		Tmp: []int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		Barriers: [][]int{},
	}

	var templates []Template

	templates = append(templates, tmp1)
	templates = append(templates, tmp2)
	templates = append(templates, tmp3)
	templates = append(templates, tmp4)
	templates = append(templates, tmp5)
	templates = append(templates, tmp6)
	// templates = append(templates, template5)
	// templates = append(templates, template6)

	// Задаем сид для рандомайзера
	rand.Seed(time.Now().UnixNano())

	blockSize := 10
	blockCount := m.SizeX / blockSize

	// Билдм мапу

	for i := 0; i < m.SizeY; i++ {
		for j := 0; j < m.SizeX; j++ {
			m.Field[i][j] = 0
		}
	}

	for i := 0; i < blockCount; i++ {
		for j := 0; j < blockCount; j++ {
			template := templates[rand.Intn(len(templates))]
			for g := 0; g < len(template.Barriers); g++ {
				bar := Barrier{}
				bar.Object = &DynamycObject{
					X:     j*blockSize*m.TileSize + template.Barriers[g][0],
					Y:     i*blockSize*m.TileSize + template.Barriers[g][1],
					Xsize: template.Barriers[g][2],
					Ysize: template.Barriers[g][3],
				}
				// bar.Object.X = template.Barriers[g][0]
				b = append(b, &bar)
				// fmt.Println("Barrier: ", template.Barriers[g])
			}
			for k := 0; k < blockSize; k++ {
				for l := 0; l < blockSize; l++ {
					m.Field[k+i*blockSize][l+j*blockSize] = template.Tmp[k*blockSize+l]
				}
			}
		}
	}

	for i := 0; i < m.SizeY; i++ {
		for j := 0; j < m.SizeX; j++ {
			// template := templates[rand.Intn(len(templates))]
			if i == 0 || i == m.SizeY-1 || j == 0 || j == m.SizeX-1 {
				m.Field[i][j] = 1
			}
		}
	}

	// Отрисовываем результат в консоль
	// for i := 0; i < 10; i++ {
	//  for j := 0; i < 10; j++ {
	//      // fmt.Println(i, j)
	//      fmt.Print(m.Field[i*10+j])
	//  }
	//  fmt.Println("")
	// }

	// for i := 0; i < 50; i++ {
	//  m.Field = append(m.Field, field[i])
	// }

	// fmt.Println(m.Field)

	return &m, b
}
