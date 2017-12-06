package model

import "testing"

func Test_getItemID(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				url: "https://qiita.com/smith_30/items/d8605f5d89cbebeae40c",
			},
			want: "d8605f5d89cbebeae40c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getItemID(tt.args.url); got != tt.want {
				t.Errorf("getItemID() = %v, want %v", got, tt.want)
			}
		})
	}
}
