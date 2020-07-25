package universe

import (
	. "testing"
)

func TestInitialUniverseShouldContainDeadEntities(t *T) {
	t.Error("TODO - Need to implement")
}

func TestInitialUniverseShouldBeGivenSize(t *T) {
	var size = 25
	var universe = InitialUniverse(size)
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
	var universe = InitialUniverse(20)

	if universe == nil {
		t.Fail()
	}
}
