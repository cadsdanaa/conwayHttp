package universe

import (
	"reflect"
	. "testing"
)

//TODO need to implement
func TestShouldSetNeighbors(t *T) {
	universeSize := 3
	//D D A
	//D D D
	//D D A
	universe := InitialUniverse(universeSize, 123)

	if universe[1][0].East() != &universe[2][0] {
		t.Error("East not set")
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
