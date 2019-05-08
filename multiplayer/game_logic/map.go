package game_logic

import (
	"fmt"
	"math/rand"
	"time"
)

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
func MapGeneration() *Map {

	var m Map

	// Инициализируем параметры карты
	m.TileSize = 10
	m.SizeX = 50
	m.SizeY = 50

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

	tmp1 := []int{
		0, 0, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 1, 1, 1, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
	}

	tmp2 := []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	tmp3 := []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 1, 1, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	tmp4 := []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 0,
		0, 0, 0, 1, 1, 1, 1, 1, 1, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	var templates [][]int

	templates = append(templates, tmp1)
	templates = append(templates, tmp2)
	templates = append(templates, tmp3)
	templates = append(templates, tmp4)
	fmt.Println(templates[0])
	// templates = append(templates, template5)
	// templates = append(templates, template6)

	// Задаем сид для рандомайзера
	rand.Seed(time.Now().UnixNano())

	blockCount := 5
	blockSize := 10
	mapSize := m.SizeX
	// templates[rand.Intn(len(templates))]
	// Билдм мапу
	for i := 0; i < 2500; i++ {
		m.Field = append(m.Field, 0)
	}
	for i := 0; i < blockCount; i++ {
		for j := 0; j < blockCount; j++ {
			template := templates[rand.Intn(len(templates))]
			for k := 0; k < blockSize; k++ {
				for l := 0; l < blockSize; l++ {
					Insert(m.Field, i*mapSize*blockSize+j*blockSize+(k*mapSize+l), template[k*blockSize+l])
				}
			}
		}
	}

	for i := 0; i < m.SizeY; i++ {
		for j := 0; j < m.SizeX; j++ {
			// template := templates[rand.Intn(len(templates))]
			if i == 0 || i == m.SizeY-1 || j == 0 || j == m.SizeX-1 {
				Insert(m.Field, i*mapSize+j, 1)
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

	fmt.Println(m.Field)
	fmt.Println("succes")

	return &m
}
