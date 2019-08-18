package conway

//Entity represents a life-form in Conway's Game of Life
type Entity struct {
	Up, Down, Left, Right *Entity
}

//UpEntity returns the reference to the neighbor above the calling entity
func (entity *Entity) UpEntity() *Entity {
	return entity.Up
}

//DownEntity returns the reference to the neighbor below the calling entity
func (entity *Entity) DownEntity() *Entity {
	return entity.Down
}

//LeftEntity returns the reference to the neighbor left of the calling entity
func (entity *Entity) LeftEntity() *Entity {
	return entity.Left
}

//RightEntity returns the reference to the neighbor right of the calling entity
func (entity *Entity) RightEntity() *Entity {
	return entity.Right
}
