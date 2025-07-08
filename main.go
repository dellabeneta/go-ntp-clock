package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/beevik/ntp"
)

// corsMiddleware permite o acesso do navegador à API.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// timeHandler agora retorna JSON em todos os casos (sucesso ou erro).
func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ntpTime, err := ntp.Time("pool.ntp.br")
	
	// Se houver um erro ao contatar o servidor NTP...
	if err != nil {
		log.Printf("ERRO: Não foi possível contatar o servidor NTP: %v", err)
		// Cria uma resposta JSON de erro.
		errorResponse := map[string]string{"error": "Falha ao sincronizar com o servidor de tempo."}
		// Define o status HTTP para Erro Interno do Servidor.
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Se tudo deu certo, cria a resposta JSON de sucesso.
	successResponse := map[string]time.Time{"time": ntpTime}
	json.NewEncoder(w).Encode(successResponse)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler)
	mux.Handle("/", http.FileServer(http.Dir("static")))

	port := "8080"
	fmt.Printf("Servidor robusto iniciado em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(mux)))
}
