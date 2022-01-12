package utils

import (
	"bytes"
	"encoding/gob"
)

func Encode(data interface{}) []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)
	if err != nil {
		HandleException(err)
	}

	return buff.Bytes()
}
