package conway

//Entity represents a life-form in Conway's Game of Life
type Entity struct {
	Up, Down, Left, Right *Entity
}
