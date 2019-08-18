package conway

import (
	"testing"
)

func TestEntityShouldHaveReferenceToNeighbors(t *testing.T) {
	var entity Entity
	if entity.Up != nil || entity.Left != nil || entity.Right != nil || entity.Down != nil {
		t.Fail()
	}
}

func TestShouldGetUpEntity(t *testing.T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.Up = &someOtherEntity

	actualEntity := someEntity.UpEntity()

	if actualEntity == nil {
		t.Error("Up entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong up entity returned")
	}
}

func TestShouldGetDownEntity(t *testing.T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.Down = &someOtherEntity

	actualEntity := someEntity.DownEntity()

	if actualEntity == nil {
		t.Error("Up entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong up entity returned")
	}
}

func TestShouldGetLeftEntity(t *testing.T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.Left = &someOtherEntity

	actualEntity := someEntity.LeftEntity()

	if actualEntity == nil {
		t.Error("Up entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong up entity returned")
	}
}

func TestShouldGetRightEntity(t *testing.T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.Right = &someOtherEntity

	actualEntity := someEntity.RightEntity()

	if actualEntity == nil {
		t.Error("Up entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong up entity returned")
	}
}
