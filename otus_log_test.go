package otus_1_6

import (
	"bytes"
	"testing"
)

func TestLogOtusEvent(t *testing.T) {
	type args struct {
		e OtusEvent
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{"homework accepted", args{HwAccepted{ID: 2, Grade: 5}}, HwAcceptedWrapper + ": ID 2, Grade 5\n"},
		{"homework submitted", args{HwSubmitted{ID: 2, Code: "some code", Comment: "some comment"}}, HwSubmittedWrapper + ":\nID: 2\nCode:\nsome code\nComment:\nsome comment\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			LogOtusEvent(tt.args.e, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("LogOtusEvent() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
