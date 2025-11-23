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
func ToString(value any) string {
	if v, ok := value.(*string); ok {
		return *v
	}
	return fmt.Sprintf("%v", value)
}

// ToBool convert any type to boolean
func ToBool(value any) bool {
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

// isInInt32Range checks if an int value is within int32 range
func isInInt32Range(value int) bool {
	return value >= int(math.MinInt32) && value <= int(math.MaxInt32)
}

// isInt64InRange checks if an int64 value is within int32 range
func isInt64InRange(value int64) bool {
	return value >= int64(math.MinInt32) && value <= int64(math.MaxInt32)
}

// isUintInRange checks if a uint value is within int32 range
func isUintInRange(value uint) bool {
	return value <= math.MaxInt32
}

// isUint32InRange checks if a uint32 value is within int32 range
func isUint32InRange(value uint32) bool {
	return value <= uint32(math.MaxInt32)
}

// isUint64InRange checks if a uint64 value is within int32 range
func isUint64InRange(value uint64) bool {
	return value <= uint64(math.MaxInt32)
}

// isFloat32InRange checks if a float32 value is within int32 range
func isFloat32InRange(value float32) bool {
	return value >= float32(math.MinInt32) && value <= float32(math.MaxInt32)
}

// isFloat64InRange checks if a float64 value is within int32 range
func isFloat64InRange(value float64) bool {
	return value >= float64(math.MinInt32) && value <= float64(math.MaxInt32)
}

// ToInt convert any type to int
func ToInt(value any) any {
	switch value := value.(type) {
	case bool:
		return boolToInt(value)
	case int:
		return intToInt(value)
	case *int:
		return ToInt(*value)
	case int8, *int8, int16, *int16, int32, *int32:
		return handleSmallInts(value)
	case int64:
		return int64ToInt(value)
	case *int64:
		return ToInt(*value)
	case uint, *uint, uint8, *uint8, uint16, *uint16:
		return handleSmallUints(value)
	case uint32:
		return uint32ToInt(value)
	case *uint32:
		return ToInt(*value)
	case uint64:
		return uint64ToInt(value)
	case *uint64:
		return ToInt(*value)
	case float32:
		return float32ToInt(value)
	case *float32:
		return ToInt(*value)
	case float64:
		return float64ToInt(value)
	case *float64:
		return ToInt(*value)
	case string:
		return stringToInt(value)
	case *string:
		return ToInt(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

// boolToInt converts bool to int
func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

// intToInt validates and converts int to int
func intToInt(value int) any {
	if !isInInt32Range(value) {
		return nil
	}
	return value
}

// int64ToInt validates and converts int64 to int
func int64ToInt(value int64) any {
	if !isInt64InRange(value) {
		return nil
	}
	return int(value)
}

// uint32ToInt validates and converts uint32 to int
func uint32ToInt(value uint32) any {
	if !isUint32InRange(value) {
		return nil
	}
	return int(value)
}

// uint64ToInt validates and converts uint64 to int
func uint64ToInt(value uint64) any {
	if !isUint64InRange(value) {
		return nil
	}
	return int(value) // #nosec G115 -- range validated above
}

// float32ToInt validates and converts float32 to int
func float32ToInt(value float32) any {
	if !isFloat32InRange(value) {
		return nil
	}
	return int(value)
}

// float64ToInt validates and converts float64 to int
func float64ToInt(value float64) any {
	if !isFloat64InRange(value) {
		return nil
	}
	return int(value)
}

// stringToInt parses string and converts to int
func stringToInt(value string) any {
	val, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil
	}
	return ToInt(val)
}

// handleSmallInts handles int8, int16, int32 and their pointers
func handleSmallInts(value any) any {
	switch v := value.(type) {
	case int8:
		return int(v)
	case *int8:
		return int(*v)
	case int16:
		return int(v)
	case *int16:
		return int(*v)
	case int32:
		return int(v)
	case *int32:
		return int(*v)
	}
	return nil
}

// handleSmallUints handles uint, uint8, uint16 and their pointers
func handleSmallUints(value any) any {
	switch v := value.(type) {
	case uint:
		if !isUintInRange(v) {
			return nil
		}
		return int(v) // #nosec G115 -- range validated above
	case *uint:
		return ToInt(*v)
	case uint8:
		return int(v)
	case *uint8:
		return int(*v)
	case uint16:
		return int(v)
	case *uint16:
		return int(*v)
	}
	return nil
}

// ToFloat convert any type to float
func ToFloat(value any) any {
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
