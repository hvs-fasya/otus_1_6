package otus_1_6

import (
	"fmt"
	"io"
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
func (hwa HwAccepted) Log() string {
	return fmt.Sprintf("%s: ID %d, Grade %d\n", HwAcceptedWrapper, hwa.ID, hwa.Grade) //example: HOMEWORK ACCEPTED: ID 2, Grade 5
}

//HwSubmitted homework submitted event struct
type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

//Log write log message
func (hws HwSubmitted) Log() string {
	return fmt.Sprintf("%s:\nID: %d\nCode:\n%s\nComment:\n%s\n", HwSubmittedWrapper, hws.ID, hws.Code, hws.Comment)
	//example: 	  HOMEWORK SUBMITTED:
	//            ID: 2
	//            Code:
	//            some code
	//            Comment:
	//            some comment
}

//OtusEvent log event interface
type OtusEvent interface {
	Log() string
}

//LogOtusEvent write e event log message to w writer
func LogOtusEvent(e OtusEvent, w io.Writer) {
	w.Write([]byte(e.Log()))
}
