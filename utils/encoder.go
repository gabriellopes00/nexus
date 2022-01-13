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
		// TODO: return the error instead of handling it
		HandleException(err)
	}

	return buff.Bytes()
}

func Decode(data []byte) interface{} {
	var decoded interface{}

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&decoded)
	if err != nil {
		HandleException(err)
	}

	return decoded
}
