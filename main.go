package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

func logRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Println("--- HTTP Request Start ---")
    fmt.Printf("Method: %s\n", r.Method)
    fmt.Printf("URL: %s\n", r.URL.String())

    fmt.Println("Headers:")
    for k, values := range r.Header {
        fmt.Printf("  %s: %s\n", k, strings.Join(values, ", "))
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        fmt.Printf("Failed to read body: %v\n", err)
    } else {
        if len(body) == 0 {
            fmt.Println("Body: <empty>")
        } else {
            fmt.Printf("Body: %s\n", string(body))
        }
    }
    fmt.Println("--- HTTP Request End ---")

    r.Body.Close()
}

func handler(w http.ResponseWriter, r *http.Request) {
    logRequest(w, r)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("OK\n"))
}

func main() {
    http.HandleFunc("/", handler)
    addr := ":8080"
    log.Printf("Listening on %s...\n", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
