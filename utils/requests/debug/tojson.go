package dbg

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToJSON(data []byte) error {
	var jsonData bytes.Buffer

	err := json.Indent(&jsonData, data, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(jsonData.String())
	return nil
}
