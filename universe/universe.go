package universe

import "github.com/cadsdanaa/conwayHttp/entity"

func InitialUniverse(universeSize int) [][]entity.Entity {
	return initialGrid(universeSize)
}

func initialGrid(universeSize int) [][]entity.Entity {
	var universe = make([][]entity.Entity, universeSize)
	for i := 0; i < universeSize; i++ {
		universe[i] = make([]entity.Entity, universeSize)
	}
	return universe
}
