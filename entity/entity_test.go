package entity

import (
	"testing"
)

func TestShouldGetUpEntity(t *testing.T) {
	var someEntity Entity
	var someOtherEntity Entity
	someEntity.SetUp(&someOtherEntity)
	actualEntity := someEntity.Up()

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
	someEntity.SetDown(&someOtherEntity)

	actualEntity := someEntity.Down()

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
	someEntity.SetLeft(&someOtherEntity)

	actualEntity := someEntity.Left()

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
	someEntity.SetRight(&someOtherEntity)

	actualEntity := someEntity.Right()

	if actualEntity == nil {
		t.Error("Up entity not assigned")
	}
	if actualEntity != &someOtherEntity {
		t.Error("Wrong up entity returned")
	}
}
