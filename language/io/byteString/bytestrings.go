package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawstring string) *bytes.Buffer {
	rawBytes := []byte(rawstring)
	var b = new(bytes.Buffer)
	b = bytes.NewBuffer(rawBytes)
	// b = bytes.NewBufferString(rawstring)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
