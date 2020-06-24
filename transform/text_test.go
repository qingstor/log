package transform

import (
	"bytes"
	"testing"
)

func TestText_Space(t *testing.T) {
	x := NewText()

	x.Key("x")
	x.AppendInt(1)

	ans := []byte("x=1")
	if !bytes.Equal(x.Bytes(), ans) {
		t.Errorf("expect %s, actual %s", ans, x.Bytes())
	}

	x.Key("y")
	x.AppendInt(2)

	ans = []byte("x=1 y=2")
	if !bytes.Equal(x.Bytes(), ans) {
		t.Errorf("expect %s, actual %s", ans, x.Bytes())
	}

	x.Key("a")
	x.Start(ContainerObject)

	ans = []byte("x=1 y=2 a={")
	if !bytes.Equal(x.Bytes(), ans) {
		t.Errorf("expect %s, actual %s", ans, x.Bytes())
	}

	x.Key("oa")
	x.AppendInt(3)

	ans = []byte("x=1 y=2 a={oa=3")
	if !bytes.Equal(x.Bytes(), ans) {
		t.Errorf("expect %s, actual %s", ans, x.Bytes())
	}

	x.Key("ob")
	x.AppendInt(4)

	ans = []byte("x=1 y=2 a={oa=3 ob=4")
	if !bytes.Equal(x.Bytes(), ans) {
		t.Errorf("expect %s, actual %s", ans, x.Bytes())
	}

	x.End(ContainerObject)

	ans = []byte("x=1 y=2 a={oa=3 ob=4}")
	if !bytes.Equal(x.Bytes(), ans) {
		t.Errorf("expect %s, actual %s", ans, x.Bytes())
	}
}
