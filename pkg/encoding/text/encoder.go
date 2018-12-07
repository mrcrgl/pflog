package text

import (
	"bytes"

	"github.com/mrcrgl/pflog/pkg/logging"

	"github.com/mrcrgl/timef"
)

type severityStringRep []byte

var (
	infoStringRep    severityStringRep = []byte("I")
	warningStringRep severityStringRep = []byte("W")
	errorStringRep   severityStringRep = []byte("E")
	fatalStringRep   severityStringRep = []byte("F")

	space                 []byte = []byte(" ")
	containerContentBegin []byte = []byte("{")
	containerContentEnd   []byte = []byte("}")
	lineEnd               []byte = []byte(";\n")
)

func Encode(in *logging.Entry) (out []byte, err error) {
	/*bs := make([]byte, 0, 26)
	b := bytes.NewBuffer(bs)*/
	b := new(bytes.Buffer)

	switch in.Severity {
	case logging.SeverityInfo:
		b.Write(infoStringRep)
		break
	case logging.SeverityWarning:
		b.Write(warningStringRep)
		break
	case logging.SeverityError:
		b.Write(errorStringRep)
		break
	case logging.SeverityFatal:
		b.Write(fatalStringRep)
		break
	}

	//fmt.Printf("len=%d\n", b.Len())

	//b.WriteString(in.Timestamp.Format(time.RFC3339))
	b.Write(timef.FormatRFC3339(in.Timestamp))

	//fmt.Printf("len=%d\n", b.Len())

	for _, c := range in.Containers {
		b.Write(space)

		if c.Enclosed() {
			b.Write(c.Kind())
			b.Write(containerContentBegin)
		}

		_, err := c.WriteTextTo(b)
		if err != nil {
			return nil, err
		}

		if c.Enclosed() {
			b.Write(containerContentEnd)
		}

		//fmt.Printf("container=%d len=%d\n", i, b.Len())
	}

	b.Write(lineEnd)

	return b.Bytes(), nil
}
