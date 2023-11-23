package utils

import (
	"fmt"
	"time"
)

func ExtractBool(data map[string]interface{}, key string) bool {
	return data[key].(bool)
}

func ExtractString(data map[string]interface{}, key string) string {
	return data[key].(string)
}

func ExtractInt32(data map[string]interface{}, key string) int32 {
	return data[key].(int32)
}

func ExtractInt64(data map[string]interface{}, key string) int64 {
	return data[key].(int64)
}

func ExtractFloat32(data map[string]interface{}, key string) float32 {
	return data[key].(float32)
}

func ExtractFloat64(data map[string]interface{}, key string) float64 {
	return data[key].(float64)
}

func ExtractArray(data map[string]interface{}, key string) []interface{} {
	return data[key].([]interface{})
}

func ExtractMap(data map[string]interface{}, key string) map[string]interface{} {
	return data[key].(map[string]interface{})
}

func ExtractTimestamp(data map[string]interface{}, key string) time.Time {
	return data[key].(time.Time)
}

func ExtractStringArray(data map[string]interface{}, key string) ([]string, error) {
	raw, ok := data[key].([]interface{})
	if !ok {
		return nil, fmt.Errorf("%s must be an array of strings", key)
	}

	var result []string
	for _, item := range raw {
		str, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("each %s must be a string", key)
		}
		result = append(result, str)
	}

	return result, nil
}

func ExtractStringArrayArray(data map[string]interface{}, key string) ([][]string, error) {
	raw, ok := data[key].([]interface{})
	if !ok {
		return nil, fmt.Errorf("%s must be an array of arrays of strings", key)
	}

	var result [][]string
	for _, item := range raw {
		strArr, ok := item.([]interface{})
		if !ok {
			return nil, fmt.Errorf("each %s must be an array of strings", key)
		}

		var strings []string
		for _, str := range strArr {
			str, ok := str.(string)
			if !ok {
				return nil, fmt.Errorf("each %s must be a string", key)
			}
			strings = append(strings, str)
		}

		result = append(result, strings)
	}

	return result, nil
}

func ExtractStringMapArray(data map[string]interface{}, key string) ([]map[string]string, error) {
	raw, ok := data[key].([]interface{})
	if !ok {
		return nil, fmt.Errorf("%s must be an array of maps of strings", key)
	}

	var result []map[string]string
	for _, item := range raw {
		strMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("each %s must be a map of strings", key)
		}

		strings := make(map[string]string)
		for k, v := range strMap {
			str, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("each %s must be a string", key)
			}
			strings[k] = str
		}

		result = append(result, strings)
	}

	return result, nil
}

func ExtractStringMap(data map[string]interface{}, key string) (map[string]string, error) {
	raw, ok := data[key].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("%s must be a map of strings", key)
	}

	result := make(map[string]string)
	for k, v := range raw {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("each %s must be a string", key)
		}
		result[k] = str
	}

	return result, nil
}
