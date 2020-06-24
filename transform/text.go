package transform

import (
	bp "github.com/Xuanwo/go-bufferpool"
	"github.com/qingstor/log/internal/pkg/bufferpool"
)

// Text is the transform both human and machine friendly.
type Text struct {
	*bp.Buffer

	keys       uint8
	containers uint8
}

// NewText will create a Transformer.
func NewText() Transformer {
	return &Text{
		Buffer: bufferpool.Get(),
	}
}

// Key will append new key.
func (t *Text) Key(key string) {
	if t.keys > t.containers {
		t.AppendByte(' ')
	} else {
		t.keys++
	}

	t.AppendString(key + "=")
}

// Start will start a container.
func (t *Text) Start(tc Container) {
	t.containers++

	switch tc {
	case ContainerObject, ContainerMap:
		t.AppendByte('{')
	case ContainerArray:
		t.AppendByte('[')
	case ContainerQuote:
		t.AppendByte('"')
	}
}

// End will end a container.
func (t *Text) End(tc Container) {
	t.containers--

	switch tc {
	case ContainerObject, ContainerMap:
		t.AppendByte('}')
	case ContainerArray:
		t.AppendByte(']')
	case ContainerQuote:
		t.AppendByte('"')
	}
}
