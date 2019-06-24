package otus_1_6

import (
	"io"
	"strconv"
)

const (
	HwAcceptedWrapper  = "HOMEWORK ACCEPTED"
	HwSubmittedWrapper = "HOMEWORK SUBMITTED"
)

//HwAccepted homework accepted event struct
type HwAccepted struct {
	ID    int
	Grade int
}

//Log write log message
func (hwa HwAccepted) Log() []byte {
	var b []byte
	b = append(b, []byte(HwAcceptedWrapper+": ID "+strconv.Itoa(hwa.ID)+", Grade "+strconv.Itoa(hwa.Grade)+"\n")...)
	return b
}

//HwSubmitted homework submitted event struct
type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

//Log write log message
func (hws HwSubmitted) Log() []byte {
	var b []byte
	b = append(b, []byte(HwSubmittedWrapper+":\nID: "+strconv.Itoa(hws.ID)+"\nCode:\n"+hws.Code+"\nComment:\n"+hws.Comment+"\n")...)
	return b
}

//OtusEvent log event interface
type OtusEvent interface {
	Log() []byte
}

//LogOtusEvent write e event log message to w writer
func LogOtusEvent(e OtusEvent, w io.Writer) {
	msg := []byte{}
	msg = append(msg, e.Log()...)
	w.Write(msg)
}
