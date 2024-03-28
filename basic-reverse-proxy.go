// A simple single-backend reverse proxy. Listens on the address given with the
// --from flag and forwards all traffic to the server given with the --to
// flag.
package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

func basicReserverProxy() {
	fromAddr := flag.String("from", "127.0.0.1:9090", "proxy's listening address")
	toAddr := flag.String("to", "127.0.0.1:8081", "the address this proxy will forward to")
	flag.Parse()

	toUrl := parseToUrl1(*toAddr)
	proxy := httputil.NewSingleHostReverseProxy(toUrl)

	log.Println("Starting proxy server on", *fromAddr)
	// if err := http.ListenAndServe(*fromAddr, proxy); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }

	// Rate limiting
	http.Handle("/", rateLimit(proxy))
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// parseToUrl parses a "to" address to url.URL value
func parseToUrl1(addr string) *url.URL {
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}
	toUrl, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}

	return toUrl
}

var limiter = rate.NewLimiter(rate.Every(5*time.Second), 2)

func rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
