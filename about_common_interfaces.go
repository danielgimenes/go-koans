package go_koans

import (
	"bytes"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		copied, err := io.Copy(out, in)
		assert(err == nil)
		assert(copied == 11)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		copied, err := io.CopyN(out, in, 5)
		assert(err == nil)
		assert(copied == 5)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
