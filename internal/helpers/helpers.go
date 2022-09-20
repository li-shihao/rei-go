package helpers

import (
	"reflect"
)

// If error panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Recursively looks for a value for a given key inside a map
func RecurseKey(m map[string]interface{}, key string) interface{} {
	for k, v := range m {
		if k == key {
			return v
		} else if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {
			if RecurseKey(v.(map[string]interface{}), key) != nil {
				return RecurseKey(v.(map[string]interface{}), key)
			}
		}
	}
	return nil
}

func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func SocketStrip(obj map[string]interface{}) map[string]interface{} {
	if _, ok := obj["certificate"]; ok {
		if _, ok := obj["certificate"].(map[string]interface{})["kind"].(map[string]interface{})["Single"]; ok {
			if k, ok := obj["certificate"].(map[string]interface{})["kind"].(map[string]interface{})["Single"].(map[string]interface{})["Call"]; ok {
				for _, s := range k.(map[string]interface{})["arguments"].([]interface{}) {
					if _, ok := s.(map[string]interface{})["Pure"]; ok {
						arr := s.(map[string]interface{})["Pure"].([]interface{})
						var arr2 []byte
						for _, j := range arr {
							arr2 = append(arr2, byte(j.(float64)))
						}
						s.(map[string]interface{})["Pure"] = string(arr2)
					}
				}
			}
		} else {
			for _, t := range obj["certificate"].(map[string]interface{})["kind"].(map[string]interface{})["Batch"].([]interface{}) {
				if k, ok := t.(map[string]interface{})["Call"]; ok {
					for _, s := range k.(map[string]interface{})["arguments"].([]interface{}) {
						if _, ok := s.(map[string]interface{})["Pure"]; ok {
							arr := s.(map[string]interface{})["Pure"].([]interface{})
							var arr2 []byte
							for _, j := range arr {
								arr2 = append(arr2, byte(j.(float64)))
							}
							s.(map[string]interface{})["Pure"] = string(arr2)
						}
					}
				}
			}
		}

		// Convert pure effect to string
		if k, ok := obj["effects"].(map[string]interface{})["events"]; ok {
			for _, s := range k.([]interface{}) {
				if v, ok := s.(map[string]interface{})["MoveEvent"]; ok {
					arr := v.(map[string]interface{})["contents"].([]interface{})
					var arr2 []byte
					for _, j := range arr {
						arr2 = append(arr2, byte(j.(float64)))
					}
					v.(map[string]interface{})["contents"] = string(arr2)
				}
			}
		}
	}

	if _, ok := obj["fields"]; ok {
		if s, ok := obj["fields"].(map[string]interface{})["Package"]; ok {
			if _map, ok := s.(map[string]interface{})["module_map"]; ok {
				for k, v := range _map.(map[string]interface{}) {
					if reflect.TypeOf(v) == reflect.TypeOf([]interface{}{}) {
						var arr []byte
						for _, j := range v.([]interface{}) {
							arr = append(arr, byte(j.(float64)))
						}
						_map.(map[string]interface{})[k] = string(arr)
					}
				}
			}
		}
	}

	return obj
}
