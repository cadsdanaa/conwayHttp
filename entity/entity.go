package entity

//Entity represents a life-form in Conway's Game of Life
type Entity struct {
	north, south, west, east, northwest, northeast, southwest, southeast *Entity
	living                                                               bool
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

func totalNeighbors(entity Entity) int {
	var neighborCount = 0
	if entity.northwest.living {
		neighborCount++
	}
	if entity.north.living {
		neighborCount++
	}
	if entity.northeast.living {
		neighborCount++
	}
	if entity.east.living {
		neighborCount++
	}
	if entity.southeast.living {
		neighborCount++
	}
	if entity.south.living {
		neighborCount++
	}
	if entity.southwest.living {
		neighborCount++
	}
	if entity.west.living {
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
