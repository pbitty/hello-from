package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	httpListenAddr := os.Getenv("HTTP_LISTEN_ADDR")
	if httpListenAddr == "" {
		httpListenAddr = ":80"
	}

	log.Printf("Listening at %s", httpListenAddr)

	err := http.ListenAndServe(httpListenAddr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.String())

		hostname, ips, err := getHostAndIps()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprintf(w, "Hello from %s (%s)", hostname, strings.Join(ips, " "))
		if err != nil {
			log.Printf("Error writing response: %s", err)
		}

	}))
	log.Println(err)
}

func getHostAndIps() (hostname string, ips []string, err error) {
	hostname, err = os.Hostname()
	if err != nil {
		return
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	ips = []string{}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipv4 := ipnet.IP.To4(); ipv4 != nil {
				ips = append(ips, ipv4.String())
			}
		}
	}
	return
}
