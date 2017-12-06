package service

import (
	"testing"
)

func Test_fetcher_fetchGridInfo(t *testing.T) {
	f := &fetcher{
		token: "7eb761c6ebadf1738b15d4b62a6a384e1a524e2f",
	}
	got, err := f.fetchGridInfo("https://qiita.com/api/v2/items/4ecc3421f2995b207284")
	if err != nil {
		t.Logf("%v\n", err)
	} else {
		t.Logf("%v\n", got.LikesCount)
	}
}
