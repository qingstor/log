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

			// set nil logger one more time
			got = ContextWithLogger(got, nil)
			logger, ok = got.Value(loggerKey).(*Logger)
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
	defaultLogger := New()
	l := &Logger{}
	ctxWithLogger := ContextWithLogger(nil, l)
	tests := []struct {
		name        string
		ctx         context.Context
		loggerExist bool
		want        *Logger
	}{
		{
			name:        "nil ctx",
			ctx:         nil,
			loggerExist: false,
			want:        defaultLogger,
		},
		{
			name:        "not set",
			ctx:         context.Background(),
			loggerExist: false,
			want:        defaultLogger,
		},
		{
			name:        "logger set",
			ctx:         ctxWithLogger,
			loggerExist: true,
			want:        l,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromContext(tt.ctx)
			if tt.loggerExist {
				if !reflect.DeepEqual(tt.want, got) {
					t.Errorf("want logger %v, got %v", tt.want, got)
				}
			} else {
				tfArg, tfRes := tt.want.transformer.(*Text), got.transformer.(*Text)
				if !reflect.DeepEqual(tfArg.vc, tfRes.vc) {
					t.Errorf("want tf %#v, got %#v", tfArg, tfRes)
				}
			}
		})
	}
}
