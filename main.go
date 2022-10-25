package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var port, timeout int
	flag.IntVar(&port, "port", 9132, "The port to listen on for HTTP requests.")
	flag.IntVar(&timeout, "timeout", 10, "Max seconds waiting for varnish report")
	flag.Parse()

	registry := prometheus.NewRegistry()
	exporter := NewPrometheusExporter(Config{
		Namespace: "cachos",
		Timeout:   time.Second * time.Duration(timeout),
		Shell:     "/bin/sh",
	})
	if err := registry.Register(exporter); err != nil {
		log.Fatalln("registry.Register failed:", err.Error())
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog: logger,
	})
	http.Handle("/metrics", handler)

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
