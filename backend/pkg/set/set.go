package set

// Set is a generic data structure that simulates a set using a map
type Set[I comparable, V any] map[I]V

// NewSet creates and returns a new instance of Set
func NewSet[I comparable, V any]() Set[I, V] {
	return make(map[I]V)
}

// Add adds a new element to the set with the given key and value
func (s *Set[I, V]) Add(k I, v V) {
	(*s)[k] = v
}

// Delete removes an element from the set based on the provided key
func (s *Set[I, V]) Delete(k I) {
	delete(*s, k)
}

// Size returns the number of elements in the set
func (s *Set[I, V]) Size() int {
	return len(*s)
}

// Exist checks if an element with the given key exists in the set
func (s *Set[I, V]) Exist(k I) bool {
	_, exists := (*s)[k]
	return exists
}

// GetByItem returns the value associated with the provided key
func (s *Set[I, V]) GetByItem(k I) V {
	return (*s)[k]
}

// GetAll returns all values stored in the set as a slice
func (s *Set[I, V]) GetAll() []V {
	slice := make([]V, 0, len(*s))
	for _, v := range *s {
		slice = append(slice, v)
	}
	return slice
}
