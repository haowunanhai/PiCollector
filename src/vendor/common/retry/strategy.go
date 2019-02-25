package retry

// Strategy This is the main interface around which this library is
// built.  It defines a very simple interface for abstracting retry
// logic in your application.
type Strategy interface {
	Next() bool
	Reset()
}
