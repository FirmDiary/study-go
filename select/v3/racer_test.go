package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
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
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn's respond within 10s", func(t *testing.T) {
		serveA := makeDelayServer(11 * time.Second)
		serveB := makeDelayServer(12 * time.Second)

		defer serveA.Close()
		defer serveB.Close()

		_, err := ConfigurableRacer(serveA.URL, serveB.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
