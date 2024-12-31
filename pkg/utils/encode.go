package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

// GobEncode encodes data using gob
func GobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
