package concurrency

import (
	"reflect"
	"testing"
	"time"
)

type WebSiteChecker func(string) bool

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Second)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100) // 切片
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
