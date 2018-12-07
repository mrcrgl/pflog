package logging

import (
	"time"
)

type Encoder func(in *Entry) (out []byte, err error)

type Decoder func(in []byte, out *Entry) (err error)

type Entry struct {
	Severity   Severity
	Timestamp  time.Time
	Containers []Container
}
