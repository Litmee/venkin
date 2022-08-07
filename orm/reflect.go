package orm

import (
	"database/sql/driver"
	"reflect"
	"time"
)

func reflectTypeObj[T interface{}](columns []string, lastcols []driver.Value) *T {

	// Database fields and values are mapped
	m := make(map[string]any)
	for i, v := range columns {
		m[v] = lastcols[i]
	}

	// receive the reflected value
	var t T

	// reflection
	pT := reflect.TypeOf(t)
	pV := reflect.ValueOf(&t)

	n := pT.NumField()

	var tag = "orm"

	for i := 0; i < n; i++ {
		// Get the name of a struct field
		// Prefer the value of the field tag orm to refer to the assignment
	reflectField:
		name := pT.Field(i).Tag.Get(tag)
		v, ok := m[name]
		if ok {
			switch v.(type) {
			case int64:
				pV.Elem().Field(i).SetInt(v.(int64))
			case []byte:
				pV.Elem().Field(i).SetString(string(v.([]byte)))
			case time.Time:
				pV.Elem().Field(i).Set(reflect.ValueOf(v))
			}
			continue
		}
		if tag == "orm" {
			// If the orm is not found then query from the tag json
			tag = "json"
			goto reflectField
		}
	}

	return &t
}
