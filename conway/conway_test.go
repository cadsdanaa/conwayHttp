package conway

import (
	"testing"
)

func TestEntityShouldHaveReferenceToNeighbors(t *testing.T) {
	var entity Entity
	if entity.Up == nil && entity.Left == nil && entity.Right == nil && entity.Down == nil {
		t.Log("Entity struct created")
	} else {
		t.Fail()
	}
}
