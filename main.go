package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      http.HandlerFunc(handler),
	}

	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(srv.Serve(ln))
}

func handler(w http.ResponseWriter, req *http.Request) {
	target := "katapult.io"
	hostname := strings.SplitN(req.Host, ":", 2)[0]

	if strings.HasSuffix(hostname, ".katapult") {
		sub := strings.TrimSuffix(hostname, ".katapult")
		switch sub {
		case "my":
			target = "my.katapult.io"
		case "io":
			target = "katapult.io"
		default:
			target = "my.katapult.io/o/" + sub
		}
	}

	w.Header().Set("Connection", "close")
	url := fmt.Sprintf("https://%s", target)
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}
