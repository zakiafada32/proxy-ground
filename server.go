package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func server1() {
	addr := flag.String("addr", "127.0.0.1:8081", "listen address")
	flag.Parse()

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			var b strings.Builder

			fmt.Fprintf(&b, "%v\t%v\t%v\tHost: %v\n", req.RemoteAddr, req.Method, req.URL, req.Host)

			for name, headers := range req.Header {
				for _, h := range headers {
					fmt.Fprintf(&b, "%v: %v\n", name, h)
				}
			}
			log.Println(b.String())
			// print all headers
			for name, headers := range req.Header {
				for _, h := range headers {
					fmt.Fprintf(&b, "%v: %v\n", name, h)
				}
			}

			fmt.Fprintf(w, "server 1: %s\n", req.URL)
		})

	log.Println("Starting server 1 on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func server2() {
	addr := flag.String("addr", "127.0.0.1:8082", "listen address")
	flag.Parse()

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			var b strings.Builder

			fmt.Fprintf(&b, "%v\t%v\t%v\tHost: %v\n", req.RemoteAddr, req.Method, req.URL, req.Host)
			for name, headers := range req.Header {
				for _, h := range headers {
					fmt.Fprintf(&b, "%v: %v\n", name, h)
				}
			}
			log.Println(b.String())

			fmt.Fprintf(w, "server 2: %s\n", req.URL)
		})

	log.Println("Starting server 2 on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
