package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Map struct {
	TileSize int   `json:"tailsize"`
	SizeX    int   `json:"sizex"`
	SizeY    int   `json:"sizey"`
	Field    []int `json:"field"`
}

// Создание карты
func MapGeneration() Map {

	var m Map

	// Инициализируем параметры карты
	m.TileSize = 10
	m.SizeX = 100
	m.SizeY = 100

	blockCount := 4

	// Логика заполнения карты препятствиями:
	// Делим карту на 16 блоков (4x4 каждый по 25 тайлов)
	// Заполняем каждый блок препятствиями

	// blockX := ( m.SizeX  ) / blockCount;
	// blockY := ( m.SizeY ) / blockCount;

	// Создаем границы карты

	// Генерируем вертикальные границы
	for i := 0; i < m.SizeY; i++ {
		for j := 0; j < m.SizeX; j++ {
			if i == 0 || i == m.SizeY-1 || j == 0 || j == m.SizeX -1 {
                m.Field[m.SizeY * i + j] = 1;
            } 
		}
	}

	// Задаем массив шаблонов карт:
	// Каждый шаблон - массив 20x20, заполенный препядствием

	var templates [][]int

	template1 := []int{
		0, 0, 1, 1, 0,
		1, 1, 1, 1, 0,
		1, 1, 1, 1, 0,
		0, 0, 1, 1, 0,
		0, 0, 1, 1, 0,
	}

	template2 := []int{
		0, 0, 0, 0, 0,
		0, 0, 1, 1, 1,
		0, 0, 1, 1, 1,
		0, 0, 1, 1, 0,
		0, 0, 1, 1, 0,
	}

	template3 := []int{
		0, 0, 0, 0, 0,
		1, 1, 0, 0, 0,
		1, 1, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 1, 1, 1, 0,
	}

	template4 := []int{
		0, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 1, 1, 1, 0,
		0, 0, 1, 1, 0,
		0, 0, 1, 1, 0,
	}

	template5 := []int{
		0, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 0, 0, 0,
		1, 1, 0, 0, 0,
	}

	template6 := []int{
		0, 0, 1, 1, 0,
		0, 0, 1, 1, 1,
		0, 0, 1, 1, 1,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
	}

	templates = append(templates, template1)
	templates = append(templates, template2)
	templates = append(templates, template3)
	templates = append(templates, template4)
	templates = append(templates, template5)
	templates = append(templates, template6)

	// Задаем сид для рандомайзера
	rand.Seed(time.Now().UnixNano())

	// Билдм мапу
	for i := 0; i < blockCount; i++ {
		for j := 0; j < blockCount; j++ {
			template := templates[rand.Intn(len(templates))]
			for k := 0; k < len(template); k++ {
				m.Field = append(m.Field, template[k])
			}
		}
	}

	fmt.Println("succes")

	// Отрисовываем результат в консоль
	// for i := 0; i < 10; i++ {
	// 	for j := 0; i < 10; j++ {
	// 		// fmt.Println(i, j)
	// 		fmt.Print(a[i*10+j])
	// 	}
	// 	fmt.Print("\n")
	// }

	return m
}

func main() {
	MapGeneration()
}
