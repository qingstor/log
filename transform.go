package log

// Container is the container type.
type Container int

// All available container type.
const (
	ContainerEntry Container = iota + 1
	ContainerObject
	ContainerMap
	ContainerArray
	ContainerQuote
)

// Transformee will implemented by Fields for transformer
type Transformee interface {
	Transform(l *Logger, e *Entry)
}

// Transformer will do transformer on a field.
type Transformer interface {
	// Key will append new key.
	Key(*Entry, string)
	// Start will pre a container.
	Start(*Entry, Container)
	// End will post a container.
	End(*Entry, Container)
}
