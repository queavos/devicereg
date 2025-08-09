package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type TokenResponse struct {
    Token    string `json:"token"`
    Hostname string `json:"hostname"`
    Status   string `json:"status"`
}

func StartLocalServer(token string) {
    hostname, err := GetHostname()
    if err != nil {
        hostname = "desconocido"
    }

    http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
        // Encabezados CORS
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // Manejo de preflight (OPTIONS)
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        response := TokenResponse{
            Token:    token,
            Hostname: hostname,
            Status:   "ok",
        }
        json.NewEncoder(w).Encode(response)
    })

    fmt.Println("Servidor local escuchando en http://localhost:8085/token")
    err = http.ListenAndServe("localhost:8085", nil)
    if err != nil {
        fmt.Println("Error iniciando servidor:", err)
    }
}
