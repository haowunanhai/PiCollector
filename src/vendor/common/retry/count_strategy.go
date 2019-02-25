package retry

// CountStrategy Try up to a fixed number of times
type CountStrategy struct {
	Tries int
	count int
}

// Next strategy
func (s *CountStrategy) Next() bool {
	s.count++
	if s.count <= s.Tries {
		return true
	}
	return false
}

// Reset strategy
func (s *CountStrategy) Reset() {
	s.count = 0
}
