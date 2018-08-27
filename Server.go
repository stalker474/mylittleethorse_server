package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

// Server Represents our api server
type Server struct {
	data *PersistObject
}

// NewServer blabla
func NewServer() (s *Server) {
	s = new(Server)
	s.data = NewPersistObject()
	return s
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Serve start the server on port port
func (s *Server) Serve(port string) error {
	http.HandleFunc("/api/json", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		data, err := s.data.toJSON()
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Fprintln(w, string(data))
		}
	})

	http.HandleFunc("/api/zjson", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		data, err := s.data.toZJSON()
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Fprintln(w, string(data))
		}
	})

	http.HandleFunc("/api/csv", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		data, err := s.data.toCSV()
		if err != nil {
			fmt.Fprintln(w, err.Error())
		} else {
			fmt.Fprintln(w, string(data))
		}
	})

	http.HandleFunc("/api/admin", func(w http.ResponseWriter, r *http.Request) {
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
			listFailedMutex.Lock()
			res, _ := json.Marshal(listFailed)
			listFailedMutex.Unlock()
			fmt.Fprintln(w, "Failed races:"+string(res))

			fmt.Fprintln(w, "Left to process: ", atomic.LoadUint64(&ops))
			return
		default:
			fmt.Fprintln(w, "Unknown method")
		}
	})

	return http.ListenAndServe(":"+port, nil)
}
