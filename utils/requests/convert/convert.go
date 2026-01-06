package convert

import (
	"encoding/json"
	"goat/utils/mapping"
	"log"
)

func ConvertTo[T any](response []byte) (*mapping.APIResponse[T], error) {
	var res mapping.APIResponse[T]
	if err := json.Unmarshal(response, &res); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &res, nil
}
