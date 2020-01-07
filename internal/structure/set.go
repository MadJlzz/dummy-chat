package structure

// Basic implementation of a dictionary that doesn't allow
// duplicates.
type Set map[string]struct{}

// Add a new element to the given set.
func (s Set) Add(value string) {
	s[value] = struct{}{}
}

// Delete an existing element from the given set.
func (s Set) Remove(value string) {
	delete(s, value)
}

// Return the size of the actual set.
func (s Set) Size() int {
	return len(s)
}

// Return whether the element exist in the set or not.
func (s Set) Has(value string) bool {
	_, ok := s[value]
	return ok
}
