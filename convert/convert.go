package convert

import (
	"fmt"
)

// ToString convert any type to string
func ToString(value interface{}) interface{} {
	if v, ok := value.(*string); ok {
		return *v
	}
	return fmt.Sprintf("%v", value)
}

// ToBool convert any type to boolean
func ToBool(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		return value
	case *bool:
		return *value
	case string:
		switch value {
		case "", "false":
			return false
		}
		return true
	case *string:
		return ToBool(*value)
	case float64:
		if value != 0 {
			return true
		}
		return false
	case *float64:
		return ToBool(*value)
	case float32:
		if value != 0 {
			return true
		}
		return false
	case *float32:
		return ToBool(*value)
	case int:
		if value != 0 {
			return true
		}
		return false
	case *int:
		return ToBool(*value)
	}
	return false
}
