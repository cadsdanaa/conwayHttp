package entity

//Entity represents a life-form in Conway's Game of Life
type Entity struct {
	north, south, west, east, northwest, northeast, southwest, southeast *Entity
	Living                                                               bool
}

//IsUnderpopulated returns true if the total number of neighbors for the entity is less than 2
func (entity Entity) IsUnderpopulated() bool {
	neighborCount := totalNeighbors(entity)
	return neighborCount < 2
}

//IsOverpopulated returns true if the total number of neighbors for the entity is more than 3
func (entity Entity) IsOverpopulated() bool {
	neighborCount := totalNeighbors(entity)
	return neighborCount > 3
}

//LivesOn returns true if the entity is currently alive and has exactly 2 or 3 neighbors
func (entity Entity) LivesOn() bool {
	neighborCount := totalNeighbors(entity)
	return entity.Living && (neighborCount == 2 || neighborCount == 3)
}

//LivesOn returns true if the entity is currently dead and has exactly 3 neighbors
func (entity Entity) Reproduces() bool {
	neighborCount := totalNeighbors(entity)
	return !entity.Living && neighborCount == 3
}

//totalNeighbors counts the number of living neighbors an entity has
func totalNeighbors(entity Entity) int {
	var neighborCount = 0
	if entity.northwest.Living {
		neighborCount++
	}
	if entity.north.Living {
		neighborCount++
	}
	if entity.northeast.Living {
		neighborCount++
	}
	if entity.east.Living {
		neighborCount++
	}
	if entity.southeast.Living {
		neighborCount++
	}
	if entity.south.Living {
		neighborCount++
	}
	if entity.southwest.Living {
		neighborCount++
	}
	if entity.west.Living {
		neighborCount++
	}
	return neighborCount
}

//North returns the reference to the neighbor north the calling entity
func (entity *Entity) North() *Entity {
	return entity.north
}

//South returns the reference to the neighbor south the calling entity
func (entity *Entity) South() *Entity {
	return entity.south
}

//West returns the reference to the neighbor west of the calling entity
func (entity *Entity) West() *Entity {
	return entity.west
}

//East returns the reference to the neighbor east of the calling entity
func (entity *Entity) East() *Entity {
	return entity.east
}

//Northwest returns the reference to the neighbor northwest of the calling entity
func (entity *Entity) Northwest() *Entity {
	return entity.northwest
}

//Northeast returns the reference to the neighbor northeast of the calling entity
func (entity *Entity) Northeast() *Entity {
	return entity.northeast
}

//Southhwest returns the reference to the neighbor southwest of the calling entity
func (entity *Entity) Southwest() *Entity {
	return entity.southwest
}

//Southeast returns the reference to the neighbor southeast of the calling entity
func (entity *Entity) Southeast() *Entity {
	return entity.southeast
}

//SetNorth sets the reference to the neighbor north of the calling entity
func (entity *Entity) SetNorth(northEntity *Entity) {
	entity.north = northEntity
}

//SetSouth sets the reference to the neighbor south of the calling entity
func (entity *Entity) SetSouth(southEntity *Entity) {
	entity.south = southEntity
}

//SetWest sets the reference to the neighbor west of the calling entity
func (entity *Entity) SetWest(westEntity *Entity) {
	entity.west = westEntity
}

//SetEast sets the reference to the neighbor east of the calling entity
func (entity *Entity) SetEast(eastEntity *Entity) {
	entity.east = eastEntity
}

//SetNorthwest sets the reference to the neighbor northwest of the calling entity
func (entity *Entity) SetNorthwest(northwestEntity *Entity) {
	entity.northwest = northwestEntity
}

//SetNortheast sets the reference to the neighbor northeast of the calling entity
func (entity *Entity) SetNortheast(northeastEntity *Entity) {
	entity.northeast = northeastEntity
}

//SetSouthwest sets the reference to the neighbor southwest of the calling entity
func (entity *Entity) SetSouthwest(southwestEntity *Entity) {
	entity.southwest = southwestEntity
}

//SetSoutheast sets the reference to the neighbor southeast of the calling entity
func (entity *Entity) SetSoutheast(southeastEntity *Entity) {
	entity.southeast = southeastEntity
}
