package wjson

import "encoding/json"

func StructToJsonString(data interface{}) string {
	dataB, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(dataB)
}
