package helpers

import "reflect"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

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

type ConnectionString struct{}

type UsernameClaim struct{}
