package universe

import (
	"github.com/cadsdanaa/conwayHttp/entity"
	"math/rand"
)

type Universe struct {
	Size     int
	Entities [][]entity.Entity
}

//InitialUniverse creates an initial universe with the given universeSize and randomizedSeed
func InitialUniverse(universeSize int, universeSeed int64) Universe {
	entities := initialGrid(universeSize)
	rand.Seed(universeSeed)
	for x, row := range entities {
		for y := range row {
			if rand.Intn(4) == 1 {
				entities[x][y] = entity.Entity{Living: true}
			}
			setNeighbors(entities, x, y)
		}
	}
	return Universe{Entities: entities, Size: universeSize}
}

func Draw(universe Universe) string {
	universeString := ""
	for x, row := range universe.Entities {
		for y := range row {
			if universe.Entities[x][y].Living {
				universeString += "@"
			} else {
				universeString += "-"
			}
		}
		universeString += "\n"
	}
	return universeString
}

//initialGrid will create a 2 dimensional array 'grid' of entities in a square of the given universeSize
func initialGrid(universeSize int) [][]entity.Entity {
	var universe = make([][]entity.Entity, universeSize)
	for i := 0; i < universeSize; i++ {
		universe[i] = make([]entity.Entity, universeSize)
	}
	return universe
}

//setNeighbors will initialize each entities neighbors in the universe
func setNeighbors(universe [][]entity.Entity, x, y int) {
	setNorthNeighbor(universe, x, y)
	setNortheastNeighbor(universe, x, y)
	setNorthwestNeighbor(universe, x, y)
	setEastNeighbor(universe, x, y)
	setWestNeighbor(universe, x, y)
	setSoutheastNeighbor(universe, x, y)
	setSouthNeighbor(universe, x, y)
	setSouthwestNeighbor(universe, x, y)
}

func setNorthNeighbor(universe [][]entity.Entity, x, y int) {
	if len(universe)-1 < (y + 1) {
		return
	}
	universe[x][y].SetNorth(&universe[x][y+1])
}

func setNortheastNeighbor(universe [][]entity.Entity, x, y int) {
	if len(universe)-1 < (y+1) || len(universe[0])-1 < (x+1) {
		return
	}
	universe[x][y].SetNortheast(&universe[x+1][y+1])
}

func setEastNeighbor(universe [][]entity.Entity, x, y int) {
	if len(universe[0])-1 < (x + 1) {
		return
	}
	universe[x][y].SetEast(&universe[x+1][y])
}

func setSoutheastNeighbor(universe [][]entity.Entity, x, y int) {
	if (y-1) < 0 || len(universe[0])-1 < (x+1) {
		return
	}
	universe[x][y].SetSoutheast(&universe[x+1][y-1])
}

func setSouthNeighbor(universe [][]entity.Entity, x, y int) {
	if (y - 1) < 0 {
		return
	}
	universe[x][y].SetSouth(&universe[x][y-1])
}

func setSouthwestNeighbor(universe [][]entity.Entity, x, y int) {
	if (y-1) < 0 || (x-1) < 0 {
		return
	}
	universe[x][y].SetSouthwest(&universe[x-1][y-1])
}

func setWestNeighbor(universe [][]entity.Entity, x, y int) {
	if (x - 1) < 0 {
		return
	}
	universe[x][y].SetWest(&universe[x-1][y])
}

func setNorthwestNeighbor(universe [][]entity.Entity, x, y int) {
	if len(universe)-1 < (y+1) || (x-1) < 0 {
		return
	}
	universe[x][y].SetNorthwest(&universe[x-1][y+1])
}
