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
