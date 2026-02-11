package xconfig

import "fmt"

type value struct {
	raw   any
	found bool
	key   string
}

func newValue(raw any, found bool, key string) Value {
	return &value{raw: raw, found: found, key: key}
}

func emptyValue(key string) Value {
	return &value{raw: nil, found: false, key: key}
}

func (v *value) String() string {
	v.mustExist()
	switch val := v.raw.(type) {
	case string:
		return val
	default:
		return fmt.Sprintf("%v", v.raw)
	}
}

func (v *value) Int() int {
	v.mustExist()
	switch val := v.raw.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	default:
		panic(fmt.Sprintf("xconfig: key %q: cannot convert %T to int", v.key, v.raw))
	}
}

func (v *value) Int64() int64 {
	v.mustExist()
	switch val := v.raw.(type) {
	case int:
		return int64(val)
	case int64:
		return val
	case float64:
		return int64(val)
	default:
		panic(fmt.Sprintf("xconfig: key %q: cannot convert %T to int64", v.key, v.raw))
	}
}

func (v *value) Float64() float64 {
	v.mustExist()
	switch val := v.raw.(type) {
	case float64:
		return val
	case int:
		return float64(val)
	case int64:
		return float64(val)
	default:
		panic(fmt.Sprintf("xconfig: key %q: cannot convert %T to float64", v.key, v.raw))
	}
}

func (v *value) Bool() bool {
	v.mustExist()
	switch val := v.raw.(type) {
	case bool:
		return val
	default:
		panic(fmt.Sprintf("xconfig: key %q: cannot convert %T to bool", v.key, v.raw))
	}
}

func (v *value) StringOr(def string) string {
	if !v.found {
		return def
	}
	switch val := v.raw.(type) {
	case string:
		return val
	default:
		return fmt.Sprintf("%v", v.raw)
	}
}

func (v *value) IntOr(def int) int {
	if !v.found {
		return def
	}
	switch val := v.raw.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	default:
		return def
	}
}

func (v *value) Int64Or(def int64) int64 {
	if !v.found {
		return def
	}
	switch val := v.raw.(type) {
	case int:
		return int64(val)
	case int64:
		return val
	case float64:
		return int64(val)
	default:
		return def
	}
}

func (v *value) Float64Or(def float64) float64 {
	if !v.found {
		return def
	}
	switch val := v.raw.(type) {
	case float64:
		return val
	case int:
		return float64(val)
	case int64:
		return float64(val)
	default:
		return def
	}
}

func (v *value) BoolOr(def bool) bool {
	if !v.found {
		return def
	}
	if val, ok := v.raw.(bool); ok {
		return val
	}
	return def
}

func (v *value) IsEmpty() bool {
	if !v.found || v.raw == nil {
		return true
	}
	if s, ok := v.raw.(string); ok && s == "" {
		return true
	}
	return false
}

func (v *value) mustExist() {
	if !v.found {
		panic(fmt.Sprintf("xconfig: key %q not found", v.key))
	}
}
