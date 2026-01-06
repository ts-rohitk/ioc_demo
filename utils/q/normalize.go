package q

import (
	"crypto/sha256"
	"encoding/hex"
	"goat/utils/mapping"
	"sort"
)

func NormalizeJson(data any) ([]byte, error) {
	jsonBytes, err := mapping.Json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var object map[string]any
	if err := mapping.Json.Unmarshal(jsonBytes, &object); err != nil {
		return nil, err
	}

	normalizedJSON := normalizeMap(object)

	return mapping.Json.Marshal(normalizedJSON)
}

func normalizeMap(data map[string]any) map[string]any {

	normalized := make(map[string]any)

	for key, value := range data {
		switch v := value.(type) {
		case map[string]any:
			normalized[key] = normalizeMap(v)
		case []any:
			normalized[key] = normalizeArray(v)
		case nil:
			normalized[key] = nil
		default:
			normalized[key] = v
		}
	}

	return normalized
}

func normalizeArray(arr []any) []any {

	if len(arr) <= 0 {
		return arr
	}

	checkAllString := true
	for _, v := range arr {
		if _, ok := v.(string); !ok {
			checkAllString = false
			break
		}
	}

	if checkAllString {
		strArrs := make([]string, 0, len(arr))
		for _, v := range arr {
			strArrs = append(strArrs, v.(string))
		}

		sort.Strings(strArrs)

		results := make([]any, 0, len(strArrs))
		for _, v := range strArrs {
			results = append(results, v)
		}

		return results
	}

	return arr
}

func rawJSONString(data mapping.RawData) (string, error) {
	raw, err := mapping.Json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

func hashNormalize(data mapping.RawData) (string, string, error) {
	bufferBytes, err := NormalizeJson(data)
	if err != nil {
		return "", "", err
	}

	hashed := sha256.Sum256(bufferBytes)
	hashedCode := hex.EncodeToString(hashed[:])
	normalizeRawJson := string(bufferBytes)

	return normalizeRawJson, hashedCode, nil
}
