package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	//use httptest we can easily create a mock HTTP server
	slowServer := makeDelayServer(20 * time.Millisecond)

	fastServer := makeDelayServer(0 * time.Millisecond)

	//在函数调用前加上defer, 该函数就会推迟在整个方法的末尾执行
	//写在这里是为了便于阅读，但是他得最后执行
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
