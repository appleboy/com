package convert

import (
	"strings"
	"testing"
)

func TestMD5Hash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hash",
			args: args{
				text: "acccityname=杭州&accname=李恪&accno=6217714100575709&accprovince=浙江&acctype=0&amount=10000&bankcode=PCBCCNBJ&currency=CNY&mhtorderno=txn20190701173504&notifyurl=https://baidu.com&opmhtid=npdown01&random=8b761ef444354229af14ed16fc3548e8&signkey=gjiowtk49Hw3l",
			},
			want: "23549444817738591679f0ceb7f77fd4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5Hash(tt.args.text); got != tt.want {
				t.Errorf("MD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

var s = strings.Repeat("s", 1024)

func BenchmarkBytesToStrOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string([]byte(s))
	}
}

func BenchmarkBytesToStrNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BytesToStr([]byte(s))
	}
}

func BenchmarkStr2BytesOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkStr2BytesNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StrToBytes(s)
	}
}

func BenchmarkConvertOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := []byte(s)
		_ = string(b)
	}
}

func BenchmarkConvertNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := StrToBytes(s)
		_ = BytesToStr(b)
	}
}
