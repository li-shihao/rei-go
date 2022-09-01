package helpers

import "reflect"

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
