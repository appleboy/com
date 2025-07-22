package convert

import (
	"encoding/binary"
	"math"
	"reflect"
	"testing"
)

func TestMD5Hash(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "ascii",
			text: "acccityname=杭州&accname=李恪&accno=6217714100575709&accprovince=浙江&acctype=0&amount=10000&bankcode=PCBCCNBJ&currency=CNY&mhtorderno=txn20190701173504&notifyurl=https://baidu.com&opmhtid=npdown01&random=8b761ef444354229af14ed16fc3548e8&signkey=gjiowtk49Hw3l",
			want: "23549444817738591679f0ceb7f77fd4",
		},
		{
			name: "unicode",
			text: "你好，世界",
			want: "dbefd3ada018615b35588a01e216ae6e",
		},
		{
			name: "empty",
			text: "",
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5Hash(tt.text); got != tt.want {
				t.Errorf("MD5Hash(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}

func TestSnakeCasedName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "Foo", "foo"},
		{"complex", "FooBarTest", "foo_bar_test"},
		{"unicode", "你好World", "你好_world"},
		{"all upper", "FOO", "f_o_o"},
		{"mixed", "fooBar你好", "foo_bar你好"},
		{"empty", "", ""},
		{"single lower", "a", "a"},
		{"single upper", "A", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SnakeCasedName(tt.in)
			if got != tt.want {
				t.Errorf("SnakeCasedName(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestTitleCasedName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"http_lib", "http_lib", "HttpLib"},
		{"id", "id", "Id"},
		{"ida", "ida", "Ida"},
		{"id_aa", "id_aa", "IdAa"},
		{"aa_id", "aa_id", "AaId"},
		{"funky", "my_r_eal_funk_ylo_ng_name", "MyREalFunkYloNgName"},
		{"unicode", "你好_world", "你好World"},
		{"unicode2", "hello_世界", "Hello世界"},
		{"empty", "", ""},
		{"single", "a", "A"},
		{"single upper", "A", "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TitleCasedName(tt.in)
			if got != tt.want {
				t.Errorf("TitleCasedName(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFloat64ToByteAndBack(t *testing.T) {
	tests := []struct {
		name string
		val  float64
	}{
		{"positive", 77.99},
		{"zero", 0},
		{"negative", -123.456},
		{"large", 1e100},
		{"small", 1e-100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Float64ToByte(tt.val)
			got := ByteToFloat64(b)
			if got != tt.val {
				t.Errorf("Float64ToByte/ByteToFloat64 roundtrip failed: got %v, want %v", got, tt.val)
			}
		})
	}
}

func TestByteToFloat64_PanicOnInvalidLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ByteToFloat64 did not panic on invalid input length")
		}
	}()
	_ = ByteToFloat64([]byte{1, 2, 3})
}

func TestFloat64ToByte_Length(t *testing.T) {
	val := 123.456
	b := Float64ToByte(val)
	if len(b) != 8 {
		t.Errorf("Float64ToByte(%v) returned %d bytes, want 8", val, len(b))
	}
}

func TestByteToFloat64_Endian(t *testing.T) {
	val := 123.456
	b := Float64ToByte(val)
	// Manually construct the expected byte slice for BigEndian
	expected := make([]byte, 8)
	binary.BigEndian.PutUint64(expected, math.Float64bits(val))
	if !reflect.DeepEqual(b, expected) {
		t.Errorf("Float64ToByte(%v) = %v, want %v", val, b, expected)
	}
}
