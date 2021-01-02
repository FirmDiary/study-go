package v1

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "zwww.cool" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"zwww.cool",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"zwww.cool":                  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
