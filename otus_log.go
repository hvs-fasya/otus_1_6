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
	b = append(b, []byte(" ID "+strconv.Itoa(hwa.ID)+", Grade "+strconv.Itoa(hwa.Grade)+"\n")...)
	return b
}

//Wrapper get log message wrapper
func (hwa HwAccepted) Wrapper() []byte {
	return []byte(HwAcceptedWrapper)
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
	b = append(b, []byte("\nID: "+strconv.Itoa(hws.ID)+"\nCode:\n"+hws.Code+"\nComment:\n"+hws.Comment+"\n")...)
	return b
}

//Wrapper get log message wrapper
func (hwa HwSubmitted) Wrapper() []byte {
	return []byte(HwSubmittedWrapper)
}

//OtusEvent log event interface
type OtusEvent interface {
	Wrapper() []byte //log with wrapper example - HOMEWORK ACCEPTED: ID 2, Grade 5
	Log() []byte
}

//LogOtusEvent write e event log message to w writer
func LogOtusEvent(e OtusEvent, w io.Writer) {
	msg := append(e.Wrapper(), []byte(":")...)
	msg = append(msg, e.Log()...)
	w.Write(msg)
}
