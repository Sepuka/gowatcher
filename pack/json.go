package pack

import (
	"bytes"
	"encoding/json"
	"io"
)

func Encode(data map[string]interface{}) io.Reader {
	out := new(bytes.Buffer)
	json.NewEncoder(out).Encode(data)

	return out
}
