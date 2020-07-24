package entity

import (
	. "testing"
)

type entityTestScenarios struct {
	inputEntity Entity
	result      bool
}

func TestShouldDetermineUnderpopulation(t *T) {
	underpopulationScenarios := []entityTestScenarios{
		{entity(nil, nil, nil, nil, nil, nil, nil, nil), true},
		{entity(&Entity{}, nil, nil, nil, nil, nil, nil, nil), true},
		{entity(nil, &Entity{}, nil, nil, nil, nil, nil, nil), true},
		{entity(nil, nil, &Entity{}, nil, nil, nil, nil, nil), true},
		{entity(nil, nil, nil, &Entity{}, nil, nil, nil, nil), true},
		{entity(nil, nil, nil, nil, &Entity{}, nil, nil, nil), true},
		{entity(nil, nil, nil, nil, nil, &Entity{}, nil, nil), true},
		{entity(nil, nil, nil, nil, nil, nil, &Entity{}, nil), true},
		{entity(nil, nil, nil, nil, nil, nil, nil, &Entity{}), true},
		{entity(&Entity{}, &Entity{}, nil, nil, nil, nil, nil, nil), false},
		{entity(nil, nil, &Entity{}, nil, nil, &Entity{}, nil, nil), false},
		{entity(nil, nil, nil, &Entity{}, nil, nil, nil, &Entity{}), false},
	}
	for _, scenario := range underpopulationScenarios {
		isUnderpopulated := scenario.inputEntity.IsUnderpopulated()
		if isUnderpopulated != scenario.result {
			t.Fail()
		}
	}
}

func TestShouldGetNorthEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetNorth(&someOtherEntity)
	actualEntity := someEntity.North()

	if actualEntity == nil {
		t.Error("North entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetSouthEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetSouth(&someOtherEntity)

	actualEntity := someEntity.South()

	if actualEntity == nil {
		t.Error("South entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetWestEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetWest(&someOtherEntity)

	actualEntity := someEntity.West()

	if actualEntity == nil {
		t.Error("West entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetEastEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetEast(&someOtherEntity)

	actualEntity := someEntity.East()

	if actualEntity == nil {
		t.Error("East entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetNorthWestEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetNorthwest(&someOtherEntity)

	actualEntity := someEntity.East()

	if actualEntity == nil {
		t.Error("Northwest entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetNorthEastEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetNortheast(&someOtherEntity)

	actualEntity := someEntity.East()

	if actualEntity == nil {
		t.Error("Northeast entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetSouthWestEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetSouthwest(&someOtherEntity)

	actualEntity := someEntity.East()

	if actualEntity == nil {
		t.Error("Southwest entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func TestShouldGetSouthEastEntity(t *T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetSoutheast(&someOtherEntity)

	actualEntity := someEntity.East()

	if actualEntity == nil {
		t.Error("Southeast entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func entity(nw, n, ne, e, se, s, sw, w *Entity) Entity {
	var testEntity Entity
	testEntity.SetNorth(n)
	testEntity.SetNorthwest(nw)
	testEntity.SetNortheast(ne)
	testEntity.SetEast(e)
	testEntity.SetSoutheast(se)
	testEntity.SetSouth(s)
	testEntity.SetSouthwest(sw)
	testEntity.SetWest(w)
	return testEntity
}
