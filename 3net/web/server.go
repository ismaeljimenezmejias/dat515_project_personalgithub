package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

// Server struct con estado (counter) y mutex para concurrencia
type Server struct {
	mu      sync.Mutex
	counter int
}

// NewServer inicializa el Server con estado inicial
func NewServer() *Server {
	return &Server{
		counter: 0,
	}
}

// ServeHTTP implementa el http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.mu.Lock()
	s.counter++
	count := s.counter
	s.mu.Unlock()

	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")

	case "/counter":
		fmt.Fprintf(w, "counter: %d\n", count)

	case "/fizzbuzz":
		valStr := r.URL.Query().Get("value")
		if valStr == "" {
			fmt.Fprintln(w, "no value provided")
			return
		}

		n, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Fprintln(w, "not an integer")
			return
		}

		fmt.Fprintln(w, fizzbuzz(n))

	case "/github":
		http.Redirect(w, r, "http://www.github.com", http.StatusMovedPermanently)

	default:
		http.NotFound(w, r)
	}
}

// lógica FizzBuzz
func fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "fizzbuzz"
	case n%3 == 0:
		return "fizz"
	case n%5 == 0:
		return "buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func main() {
	flag.Parse()
	server := NewServer()
	log.Println("Server running on", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, server))
}
