package internal

type Map struct {
	values map[Coordinates]Entity
	Height uint8
	Width  uint8
}

type Coordinates struct {
	X, Y uint8
}

func NewMap(height uint8, width uint8) *Map {
	return &Map{
		values: map[Coordinates]Entity{},
		Height: height,
		Width:  width,
	}
}

func (m Map) IsEmpty(c Coordinates) bool {
	_, occupied := m.values[c]
	return !occupied
}

func (m Map) IsValid(c Coordinates) bool {
	if c.X >= m.Width {
		return false
	}
	if c.Y >= m.Height {
		return false
	}
	if c.X < 0 || c.Y < 0 {
		return false
	}
	return true
}

func (m *Map) Move(from Coordinates, to Coordinates) {
	entity := m.values[from]
	delete(m.values, from)
	m.values[to] = entity
}

func (m *Map) Delete(from Coordinates) {
	delete(m.values, from)
}

func Find[T Entity](m Map, coordinates Coordinates) (T, bool) {
	entity, found := m.values[coordinates]

	if found {
		if e, ok := entity.(T); ok {
			return e, found
		}
	}

	var zeroValue T
	return zeroValue, found
}

func FindByType[T Entity](m Map) []T {
	entities := []T{}

	for _, v := range m.values {
		if e, ok := v.(T); ok {
			entities = append(entities, e)
		}
	}

	return entities
}

func (m *Map) Add(entity Entity, coordinates Coordinates) {
	m.values[coordinates] = entity
}
