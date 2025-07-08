package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/beevik/ntp"
)

var (
	lastNtpTime   time.Time
	lastNtpFetch  time.Time
	cacheDuration = 30 * time.Second
	cacheMutex    sync.Mutex
)

// corsMiddleware permite o acesso do navegador à API.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// timeHandler retorna a hora NTP em cache, atualizando se necessário.
func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	now := time.Now()
	if lastNtpTime.IsZero() || now.Sub(lastNtpFetch) > cacheDuration {
		ntpTime, err := ntp.Time("pool.ntp.br")
		if err != nil {
			log.Printf("ERRO: Não foi possível contatar o servidor NTP: %v", err)
			errorResponse := map[string]string{"error": "Falha ao sincronizar com o servidor de tempo."}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		lastNtpTime = ntpTime
		lastNtpFetch = now
	}

	successResponse := map[string]time.Time{"time": lastNtpTime.Add(now.Sub(lastNtpFetch))}
	json.NewEncoder(w).Encode(successResponse)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler)
	mux.Handle("/", http.FileServer(http.Dir("static")))

	port := "8080"
	fmt.Printf("Servidor iniciado em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(mux)))
}
