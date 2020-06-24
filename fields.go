package log

import (
	"time"

	"github.com/qingstor/log/transform"
)

// IntField carries an int value.
type IntField struct {
	k string
	v int64
}

// Int creates an IntField
func Int(k string, v int64) *IntField {
	return &IntField{
		k: k,
		v: v,
	}
}

// Transform an IntField.
func (f *IntField) Transform(m transform.Transformer) {
	m.Key(f.k)
	m.AppendInt(f.v)
}

// UintField carries a uint value.
type UintField struct {
	k string
	v uint64
}

// Uint creates a UintField
func Uint(k string, v uint64) *UintField {
	return &UintField{
		k: k,
		v: v,
	}
}

// Transform a UintField.
func (f *UintField) Transform(m transform.Transformer) {
	m.Key(f.k)
	m.AppendUint(f.v)
}

// StringField carries a string value.
type StringField struct {
	k string
	v string
}

// String creates a StringField
func String(k, v string) *StringField {
	return &StringField{
		k: k,
		v: v,
	}
}

// Transform a StringField.
func (f *StringField) Transform(m transform.Transformer) {
	m.Key(f.k)

	m.Start(transform.ContainerQuote)
	defer m.End(transform.ContainerQuote)

	m.AppendString(f.v)
}

// BytesField carries an bytes value.
type BytesField struct {
	k string
	v []byte
}

// Bytes creates an BytesField
func Bytes(k string, v []byte) *BytesField {
	return &BytesField{
		k: k,
		v: v,
	}
}

// Transform a BytesField.
func (f *BytesField) Transform(m transform.Transformer) {
	m.Key(f.k)
	m.AppendBytes(f.v)
}

// FloatField carries a float value.
type FloatField struct {
	k string
	v float64
}

// Float creates an FloatField
func Float(k string, v float64) *FloatField {
	return &FloatField{
		k: k,
		v: v,
	}
}

// Transform a FloatField.
func (f *FloatField) Transform(m transform.Transformer) {
	m.Key(f.k)
	m.AppendFloat(f.v)
}

// TimeField carries a time value.
type TimeField struct {
	k string
	v time.Time

	layout string
}

// Time creates an TimeField
func Time(k string, v time.Time, layout string) *TimeField {
	return &TimeField{
		k: k,
		v: v,

		layout: layout,
	}
}

// Transform a TimeField.
func (f *TimeField) Transform(m transform.Transformer) {
	m.Key(f.k)
	m.AppendTime(f.v, f.layout)
}
