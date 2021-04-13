package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubWebsiteChecker(string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

///////////////////////////
func mockWebsiteChecker(u string) bool {
	if u == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestWebsiteChecker(t *testing.T) {

	websites := []string{
		"http://www.google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://www.google.com":      true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, got:%v", want, got)
	}
}
