package log

import (
	"context"
	"reflect"
	"testing"
)

func TestContextWithLogger(t *testing.T) {
	l := &Logger{}

	tests := []struct {
		name string
		ctx  context.Context
		l    *Logger
	}{
		{
			name: "nil ctx",
			ctx:  nil,
			l:    l,
		},
		{
			name: "todo ctx",
			ctx:  context.TODO(),
			l:    l,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContextWithLogger(tt.ctx, tt.l)
			logger, ok := got.Value(loggerKey).(*Logger)
			if !ok {
				t.Errorf("not validated logger")
			}
			if !reflect.DeepEqual(tt.l, logger) {
				t.Errorf("want %v, got %v", tt.l, logger)
			}
		})
	}
}

func TestFromContext(t *testing.T) {
	l := &Logger{}
	ctxWithLogger := ContextWithLogger(nil, l)
	tests := []struct {
		name string
		ctx  context.Context
		want *Logger
	}{
		{
			name: "nil ctx",
			ctx:  nil,
			want: nil,
		},
		{
			name: "not set",
			ctx:  context.Background(),
			want: nil,
		},
		{
			name: "logger set",
			ctx:  ctxWithLogger,
			want: l,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromContext(tt.ctx)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}
