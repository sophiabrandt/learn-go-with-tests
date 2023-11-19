package race

import (
	"fmt"
	"net/http"
	"time"
)

// URLPing is the result of a ping, including the URL and any error encountered
type URLPing struct {
	URL string
	Err error
}

var tenSecondTimeout = 10 * time.Second

// Racer checks for a fast url
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer checks for a fast url with a configurable timeout
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case result := <-ping(a):
		if result.Err != nil {
			return "", result.Err
		}
		return result.URL, nil
	case result := <-ping(b):
		if result.Err != nil {
			return "", result.Err
		}
		return result.URL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for urls to respond")
	}
}

func ping(url string) chan URLPing {
	ch := make(chan URLPing)
	go func() {
		_, err := http.Get(url)
		ch <- URLPing{URL: url, Err: err}
		close(ch)
	}()
	return ch
}
