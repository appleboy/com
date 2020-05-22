package convert

import (
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

func TestSnakeCasedName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				name: "Foo",
			},
			want: "foo",
		},
		{
			name: "complex",
			args: args{
				name: "FooBarTest",
			},
			want: "foo_bar_test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCasedName(tt.args.name); got != tt.want {
				t.Errorf("SnakeCasedName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitleCasedName(t *testing.T) {
	testCases := map[string]string{
		"http_lib":                  "HttpLib",
		"id":                        "Id",
		"ida":                       "Ida",
		"id_aa":                     "IdAa",
		"aa_id":                     "AaId",
		"my_r_eal_funk_ylo_ng_name": "MyREalFunkYloNgName",
	}

	for in, expected := range testCases {
		out := TitleCasedName(in)
		if out != expected {
			t.Errorf("Given %s, expected %s but got %s", in, expected, out)
		}
	}
}
