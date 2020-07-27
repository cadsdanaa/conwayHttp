package universe

import (
	"github.com/cadsdanaa/conwayHttp/entity"
	"math/rand"
)

//InitialUniverse creates an initial universe with the given universeSize and randomizedSeed
func InitialUniverse(universeSize int, universeSeed int64) [][]entity.Entity {
	universe := initialGrid(universeSize)
	rand.Seed(universeSeed)
	for r, row := range universe {
		for i := range row {
			if rand.Intn(4) == 1 {
				universe[r][i] = entity.Entity{Living: true}
			}
		}
	}
	return universe
}

//initialGrid will create a 2 dimensional array 'grid' of entities in a square of the given universeSize
func initialGrid(universeSize int) [][]entity.Entity {
	var universe = make([][]entity.Entity, universeSize)
	for i := 0; i < universeSize; i++ {
		universe[i] = make([]entity.Entity, universeSize)
	}
	return universe
}
