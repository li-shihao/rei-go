package main

import "reflect"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func recurseKey(m map[string]interface{}, key string) interface{} {
	for k, v := range m {
		if k == key {
			return v
		} else if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {
			if recurseKey(v.(map[string]interface{}), key) != nil {
				return recurseKey(v.(map[string]interface{}), key)
			}
		}
	}
	return nil
}
