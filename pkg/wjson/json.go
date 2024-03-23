package wjson

import "encoding/json"

func StructToJsonString(data interface{}) string {
	dataB, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(dataB)
}

func StructToJsonStringWithIndent(data interface{}, prefix, indent string) string {
	dataB, err := json.MarshalIndent(data, prefix, indent)
	if err != nil {
		return ""
	}
	return string(dataB)
}
