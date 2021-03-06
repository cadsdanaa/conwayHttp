package entity

import (
	. "testing"
)

type entityTestScenarios struct {
	inputEntity Entity
	result      bool
}

func TestShouldReproduce(t *T) {
	liveOnScenarios := []entityTestScenarios{
		{deadEntity(dead(), dead(), dead(), dead(), dead(), dead(), dead(), dead()), false},
		{deadEntity(alive(), dead(), dead(), dead(), dead(), dead(), dead(), dead()), false},
		{deadEntity(alive(), alive(), dead(), dead(), dead(), dead(), dead(), dead()), false},
		{deadEntity(alive(), dead(), dead(), dead(), dead(), dead(), dead(), alive()), false},
		{deadEntity(alive(), alive(), alive(), dead(), dead(), dead(), dead(), dead()), true},
		{deadEntity(alive(), dead(), dead(), dead(), alive(), dead(), dead(), alive()), true},
	}
	for _, scenario := range liveOnScenarios {
		reproduces := scenario.inputEntity.Reproduces()
		if reproduces != scenario.result {
			t.Fail()
		}
	}
}

func TestShouldLiveOn(t *T) {
	liveOnScenarios := []entityTestScenarios{
		{entity(dead(), dead(), dead(), dead(), dead(), dead(), dead(), dead()), false},
		{entity(alive(), dead(), dead(), dead(), dead(), dead(), dead(), dead()), false},
		{deadEntity(alive(), alive(), alive(), dead(), dead(), dead(), dead(), dead()), false},
		{entity(alive(), alive(), dead(), dead(), dead(), dead(), dead(), dead()), true},
		{entity(alive(), alive(), alive(), dead(), dead(), dead(), dead(), dead()), true},
		{entity(alive(), dead(), dead(), dead(), dead(), dead(), dead(), alive()), true},
		{entity(alive(), dead(), dead(), dead(), alive(), dead(), dead(), alive()), true},
	}
	for _, scenario := range liveOnScenarios {
		livesOn := scenario.inputEntity.LivesOn()
		if livesOn != scenario.result {
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

	actualEntity := someEntity.Northwest()

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

	actualEntity := someEntity.Northeast()

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

	actualEntity := someEntity.Southwest()

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

	actualEntity := someEntity.Southeast()

	if actualEntity == nil {
		t.Error("Southeast entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong entity returned")
	}
}

func entity(nw, n, ne, e, se, s, sw, w *Entity) Entity {
	var testEntity Entity
	testEntity.Living = true
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

func deadEntity(nw, n, ne, e, se, s, sw, w *Entity) Entity {
	testEntity := entity(nw, n, ne, e, se, s, sw, w)
	testEntity.Living = false
	return testEntity
}

func alive() *Entity {
	return &Entity{Living: true}
}

func dead() *Entity {
	return &Entity{}
}
