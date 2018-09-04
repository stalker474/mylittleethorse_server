package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

// Server Represents our api server
type Server struct {
	data  *PersistObject
	cache map[string]string
}

// NewServer blabla
func NewServer() (s *Server) {
	s = new(Server)
	s.data = NewPersistObject()
	s.resetCache()
	return s
}

func (s *Server) resetCache() {
	s.cache = make(map[string]string)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableDecorators(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/plain")
}

func enableDecoratorsGz(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Encoding", "gzip")
	(*w).Header().Set("Content-Type", "text/javascript")
}

func getFromAndTo(r *http.Request) (from uint32, to uint32, err error) {
	keysFrom, okFrom := r.URL.Query()["from"]
	keysTo, okTo := r.URL.Query()["to"]

	from = 0
	to = 9999

	if okFrom && (len(keysFrom) > 0) {
		val, err := strconv.ParseInt(keysFrom[0], 10, 32)
		if err != nil {
			return 0, 0, err
		}

		from = uint32(val)
	}
	if okTo && (len(keysTo) > 0) {
		val, err := strconv.ParseInt(keysTo[0], 10, 32)
		if err != nil {
			return 0, 0, err
		}
		to = uint32(val)
	}

	return from, to, nil
}

// Serve start the server on port port
func (s *Server) Serve(port string) error {
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		enableDecorators(&w)
		enableCors(&w)
		from, to, err := getFromAndTo(r)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		req := "json" + strconv.Itoa(int(from)) + "_" + strconv.Itoa(int(to))
		_, exists := s.cache[req]
		if exists {
			fmt.Fprintln(w, s.cache[req])
		} else {
			data, err := s.data.toJSON(from, to)
			if err != nil {
				fmt.Fprintln(w, err.Error())
			} else {
				str := string(data[:])
				fmt.Fprintln(w, str)
				s.cache[req] = str
			}
		}
	})

	http.HandleFunc("/bridge", func(w http.ResponseWriter, r *http.Request) {
		enableDecorators(&w)
		enableCors(&w)
		req := "bridge"
		_, exists := s.cache[req]
		if exists {
			fmt.Fprintln(w, s.cache[req])
		} else {
			data, err := s.data.toLightJSON()
			if err != nil {
				fmt.Fprintln(w, err.Error())
			} else {
				str := string(data[:])
				fmt.Fprintln(w, str)
				s.cache[req] = str
			}
		}
	})

	http.HandleFunc("/zjson", func(w http.ResponseWriter, r *http.Request) {
		enableDecorators(&w)
		enableCors(&w)
		from, to, err := getFromAndTo(r)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		req := "zjson" + strconv.Itoa(int(from)) + "_" + strconv.Itoa(int(to))
		_, exists := s.cache[req]
		if exists {
			fmt.Fprintln(w, s.cache[req])
		} else {
			data, err := s.data.toZJSON(from, to)
			if err != nil {
				fmt.Fprintln(w, err.Error())
			} else {
				str := string(data[:])
				fmt.Fprintln(w, str)
				s.cache[req] = str
			}
		}
	})

	http.HandleFunc("/csv", func(w http.ResponseWriter, r *http.Request) {
		enableDecorators(&w)
		enableCors(&w)
		from, to, err := getFromAndTo(r)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		req := "zjson" + strconv.Itoa(int(from)) + "_" + strconv.Itoa(int(to))
		_, exists := s.cache[req]
		if exists {
			fmt.Fprintln(w, s.cache[req])
		} else {
			data, err := s.data.toCSV(from, to)
			if err != nil {
				fmt.Fprintln(w, err.Error())
			} else {
				str := string(data[:])
				fmt.Fprintln(w, str)
				s.cache[req] = str
			}
		}
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		enableDecorators(&w)
		enableCors(&w)
		keys, ok := r.URL.Query()["method"]

		if !ok || len(keys[0]) < 1 {
			fmt.Fprintln(w, "Missing method")
			fmt.Fprintln(w, "updatedb : refreshes json cache data from current state")
			fmt.Fprintln(w, "recache : refresh ALL data keeping only old ethorse bridge info")
			fmt.Fprintln(w, "report : gives info about current update status")
			return
		}

		switch keys[0] {
		case "updatedb":
			s.data.save()
			fmt.Fprintln(w, "Recached")
			return
		case "recache":
			atomic.StoreUint32(&fullRefresh, 1)
			fmt.Fprintln(w, "Full recache ordered")
			return
		case "report":
			fmt.Fprintln(w, "Left to process: ", atomic.LoadUint64(&ops))
			return
		default:
			fmt.Fprintln(w, "Unknown method")
		}
	})

	//return http.ListenAndServeTLS(":443", "server.rsa.crt", "server.rsa.key", nil)

	return http.ListenAndServe(":"+port, nil)
}
