package transform

import (
	"time"
)

// Container is the container type.
type Container int

// All available container type.
const (
	ContainerObject Container = iota + 1
	ContainerMap
	ContainerArray
	ContainerQuote
)

// FactoryFunc is func used for transformer factory
type FactoryFunc func() Transformer

// Transformee will implemented by Fields for transform
type Transformee interface {
	Transform(m Transformer)
}

// Transformer will do transform on a field.
type Transformer interface {
	// Bytes will return all bytes in the transform.
	Bytes() []byte
	// Free will free a transformer.
	Free()

	// Key will append new key.
	Key(string)
	// Start will start a container.
	Start(Container)
	// End will end a container.
	End(Container)

	AppendBool(v bool)
	AppendByte(v byte)
	AppendBytes(v []byte)
	AppendFloat(f float64)
	AppendInt(i int64)
	AppendString(s string)
	AppendTime(t time.Time, layout string)
	AppendUint(i uint64)
}
