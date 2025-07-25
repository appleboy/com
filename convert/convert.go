package convert

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/appleboy/com/bytesconv"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// ToString convert any type to string
func ToString(value interface{}) string {
	if v, ok := value.(*string); ok {
		return *v
	}
	return fmt.Sprintf("%v", value)
}

// ToBool convert any type to boolean
func ToBool(value interface{}) bool {
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

// ToInt convert any type to int
func ToInt(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value {
			return 1
		}
		return 0
	case int:
		if value < int(math.MinInt32) || value > int(math.MaxInt32) {
			return nil
		}
		return value
	case *int:
		return ToInt(*value)
	case int8:
		return int(value)
	case *int8:
		return int(*value)
	case int16:
		return int(value)
	case *int16:
		return int(*value)
	case int32:
		return int(value)
	case *int32:
		return int(*value)
	case int64:
		if value < int64(math.MinInt32) || value > int64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *int64:
		return ToInt(*value)
	case uint:
		if value > math.MaxInt32 {
			return nil
		}
		return int(value)
	case *uint:
		return ToInt(*value)
	case uint8:
		return int(value)
	case *uint8:
		return int(*value)
	case uint16:
		return int(value)
	case *uint16:
		return int(*value)
	case uint32:
		if value > uint32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *uint32:
		return ToInt(*value)
	case uint64:
		if value > uint64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *uint64:
		return ToInt(*value)
	case float32:
		if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *float32:
		return ToInt(*value)
	case float64:
		if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *float64:
		return ToInt(*value)
	case string:
		val, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return nil
		}
		return ToInt(val)
	case *string:
		return ToInt(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

// ToFloat convert any type to float
func ToFloat(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value {
			return 1.0
		}
		return 0.0
	case *bool:
		return ToFloat(*value)
	case int:
		return float64(value)
	case *int32:
		return ToFloat(*value)
	case float32:
		return value
	case *float32:
		return ToFloat(*value)
	case float64:
		return value
	case *float64:
		return ToFloat(*value)
	case string:
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil
		}
		return val
	case *string:
		return ToFloat(*value)
	}
	return 0.0
}

// ToPtr takes a value of any type and returns a pointer to that value.
// This is useful for converting a value to a pointer in a generic way.
//
// T: The type of the value being converted.
// value: The value to be converted to a pointer.
//
// Returns a pointer to the provided value.
func ToPtr[T any](value T) *T {
	return &value
}

// FromPtr takes a pointer to a value and returns the value itself.
// If the pointer is nil, it returns the zero value of the type.
//
// T: The type of the value being dereferenced.
// ptr: The pointer to the value.
//
// Returns the value pointed to by the pointer, or the zero value if the pointer is nil.
func FromPtr[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// ConvertBig5ToUTF8 converts a string encoded in Big5 to a UTF-8 encoded string.
// The input parameter s must be a string encoded in Big5, and the function returns the corresponding UTF-8 string.
// This function uses the Go standard library's transform package for conversion, which results in a string allocation.
// If the conversion fails (e.g., if the input is not valid Big5), the original string is returned and no panic occurs.
//
// Usage Example:
//
//	big5Str := "\xa4\xa4\xa4\xe5" // "中文" in Big5 encoding
//	utf8Str := ConvertBig5ToUTF8(big5Str)
//	fmt.Println(utf8Str) // Output: 中文
func ConvertBig5ToUTF8(s string) string {
	reader := transform.NewReader(
		strings.NewReader(s),
		traditionalchinese.Big5.NewDecoder(),
	)
	d, err := io.ReadAll(reader)
	if err != nil {
		return s
	}
	return bytesconv.BytesToStr(d)
}
