package internal

import (
	"fmt"
)

type PathFinder[T Entity] interface {
	FindPath(Map, Coordinates) []Coordinates
}

type BFSPathFinder[T Entity] struct {
	target T
}

func NewBFSPathFinder[T Entity]() BFSPathFinder[T] {
	return BFSPathFinder[T]{}
}

func (bfspf BFSPathFinder[T]) FindPath(m Map, origin Coordinates) []Coordinates {
	q := queue{}
	visited := newSet()
	predecessors := make(map[Coordinates]Coordinates)

	q.enqueue(origin)

	for !q.isEmpty() {
		current, err := q.dequeue()
		if err != nil {
			break
		}

		if bfspf.isTarget(m, current) {
			return reconstrucPath(predecessors, current)
		}

		visited.add(current)
		for _, neighbour := range findNeighbours(m, current) {
			if q.contains(neighbour) || visited.contains(neighbour) {
				continue
			}
			if isObstacle(m, neighbour) && !bfspf.isTarget(m, neighbour) {
				continue
			}

			q.enqueue(neighbour)
			predecessors[neighbour] = current
		}
	}

	return []Coordinates{}
}

func findNeighbours(m Map, origin Coordinates) []Coordinates {
	var neighbours []Coordinates

	for x := int(origin.X) - 1; x <= int(origin.X)+1; x++ {
		for y := int(origin.Y) - 1; y <= int(origin.Y)+1; y++ {
			neighbourCoordinates := Coordinates{X: uint8(x), Y: uint8(y)}

			if origin == neighbourCoordinates {
				continue
			}

			if m.IsValid(neighbourCoordinates) {
				neighbours = append(neighbours, neighbourCoordinates)
			}
		}
	}

	return neighbours
}

func reconstrucPath(predecessors map[Coordinates]Coordinates, last Coordinates) []Coordinates {
	var path []Coordinates

	path = append(path, last)

	for {
		if _, found := predecessors[last]; !found {
			break
		}
		last = predecessors[last]
		path = append(path, last)
	}

	path = path[:len(path)-1]

	return reverse(path)
}

func reverse(s []Coordinates) []Coordinates {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return s
}

func isObstacle(m Map, c Coordinates) bool {
	if _, found := Find[Entity](m, c); found {
		return true
	}
	return false
}

func (bfspf BFSPathFinder[T]) isTarget(m Map, c Coordinates) bool {
	entity, found := Find[Entity](m, c)

	if found {
		if _, ok := entity.(T); ok {
			return true
		}
	}

	return false
}

type queue struct {
	values []Coordinates
}

func (q *queue) enqueue(c Coordinates) {
	q.values = append(q.values, c)
}

func (q *queue) dequeue() (Coordinates, error) {
	if q.isEmpty() {
		return Coordinates{}, fmt.Errorf("queue is empty")
	}

	first := q.values[0]
	q.values = q.values[1:]

	return first, nil
}

func (q queue) isEmpty() bool {
	if len(q.values) == 0 {
		return true
	}
	return false
}

func (q queue) contains(c Coordinates) bool {
	for _, value := range q.values {
		if value == c {
			return true
		}
	}
	return false
}

type set struct {
	values map[Coordinates]bool
}

func newSet() *set {
	return &set{values: make(map[Coordinates]bool)}
}

func (s *set) add(c Coordinates) {
	s.values[c] = true
}

func (s *set) contains(c Coordinates) bool {
	return s.values[c]
}
