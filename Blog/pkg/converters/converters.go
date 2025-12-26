package converters

import "encoding/json"

func MapToString(mapData map[string]string) string {
	str, _ := json.Marshal(mapData)
	return string(str)
}

func StringToMap(strData string) map[string]string {
	var out map[string]string
	_ = json.Unmarshal([]byte(strData), &out)
	return out
}
