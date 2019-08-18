package conway

//Entity represents a life-form in Conway's Game of Life
type Entity struct {
	up, down, left, right *Entity
}

//Up returns the reference to the neighbor above the calling entity
func (entity *Entity) Up() *Entity {
	return entity.up
}

//Down returns the reference to the neighbor below the calling entity
func (entity *Entity) Down() *Entity {
	return entity.down
}

//Left returns the reference to the neighbor left of the calling entity
func (entity *Entity) Left() *Entity {
	return entity.left
}

//Right returns the reference to the neighbor right of the calling entity
func (entity *Entity) Right() *Entity {
	return entity.right
}

//SetUp sets the reference to the neighbor above the calling entity
func (entity *Entity) SetUp(upEntity *Entity) {
	entity.up = upEntity
}

//SetDown sets the reference to the neighbor below the calling entity
func (entity *Entity) SetDown(downEntity *Entity) {
	entity.down = downEntity
}

//SetLeft sets the reference to the neighbor below the calling entity
func (entity *Entity) SetLeft(leftEntity *Entity) {
	entity.left = leftEntity
}

//SetRight sets the reference to the neighbor below the calling entity
func (entity *Entity) SetRight(rightEntity *Entity) {
	entity.right = rightEntity
}
