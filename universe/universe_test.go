package universe

import (
	"github.com/cadsdanaa/conwayHttp/entity"
	"reflect"
	"strings"
	. "testing"
)

func TestShouldDrawUniverse(t *T) {
	expected := "-@\n--\n"
	universe := InitialUniverse(2, 123)

	actual := Draw(universe)

	if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
		t.Fail()
	}
}

func TestShouldSetNeighbors(t *T) {
	universeSize := 3
	universe := InitialUniverse(universeSize, 123)

	for r, row := range universe {
		for i := range row {
			checkNorth(universe, i, r, t)
			checkNortheast(universe, i, r, t)
			checkNorthwest(universe, i, r, t)
			checkEast(universe, i, r, t)
			checkWest(universe, i, r, t)
			checkSoutheast(universe, i, r, t)
			checkSouth(universe, i, r, t)
			checkSouthwest(universe, i, r, t)
		}
	}

}

func checkNorth(universe [][]entity.Entity, x, y int, t *T) {
	if len(universe)-1 < (y + 1) {
		return
	}
	if universe[x][y].North() != &universe[x][y+1] {
		t.Error("North not set correctly")
	}
}

func checkNortheast(universe [][]entity.Entity, x, y int, t *T) {
	if len(universe)-1 < (y+1) || len(universe[0])-1 < (x+1) {
		return
	}
	if universe[x][y].Northeast() != &universe[x+1][y+1] {
		t.Error("Northeast not set correctly")
	}
}

func checkEast(universe [][]entity.Entity, x, y int, t *T) {
	if len(universe[0])-1 < (x + 1) {
		return
	}
	if universe[x][y].East() != &universe[x+1][y] {
		t.Error("East not set correctly")
	}
}

func checkSoutheast(universe [][]entity.Entity, x, y int, t *T) {
	if (y-1) < 0 || len(universe[0])-1 < (x+1) {
		return
	}
	if universe[x][y].Southeast() != &universe[x+1][y-1] {
		t.Error("Southeast not set correctly")
	}
}

func checkSouth(universe [][]entity.Entity, x, y int, t *T) {
	if (y - 1) < 0 {
		return
	}
	if universe[x][y].South() != &universe[x][y-1] {
		t.Error("South not set correctly")
	}
}

func checkSouthwest(universe [][]entity.Entity, x, y int, t *T) {
	if (y-1) < 0 || (x-1) < 0 {
		return
	}
	if universe[x][y].Southwest() != &universe[x-1][y-1] {
		t.Error("Southwest not set correctly")
	}
}

func checkWest(universe [][]entity.Entity, x, y int, t *T) {
	if (x - 1) < 0 {
		return
	}
	if universe[x][y].West() != &universe[x-1][y] {
		t.Error("West not set correctly")
	}
}

func checkNorthwest(universe [][]entity.Entity, x, y int, t *T) {
	if len(universe)-1 < (y+1) || (x-1) < 0 {
		return
	}
	if universe[x][y].Northwest() != &universe[x-1][y+1] {
		t.Error("Northwest not set correctly")
	}
}

func TestShouldCreateADeterministicallySeededUniverse(t *T) {
	universe1 := InitialUniverse(10, 123)
	universe2 := InitialUniverse(10, 123)
	universe3 := InitialUniverse(10, 0)
	if !reflect.DeepEqual(universe1, universe2) {
		t.Error("universe1 and universe2 had the same seed thus they should be equal")
	}
	if reflect.DeepEqual(universe1, universe3) {
		t.Error("universe1 should have living entities")
	}
}

func TestInitialUniverseShouldBeGivenSize(t *T) {
	size := 25
	universe := InitialUniverse(size, 0)
	if len(universe) != size {
		t.Fail()
	}
	for _, row := range universe {
		if len(row) != size {
			t.Fail()
		}
	}
}

func TestShouldInitializeUniverse(t *T) {
	universe := InitialUniverse(20, 0)

	if universe == nil {
		t.Fail()
	}
}
