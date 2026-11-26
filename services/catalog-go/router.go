package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Product is one item in the in-memory catalog.
type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PriceCents int    `json:"priceCents"`
}

var products = []Product{
	{ID: "p-1", Name: "Notebook", PriceCents: 450},
	{ID: "p-2", Name: "Pen", PriceCents: 120},
	{ID: "p-3", Name: "Backpack", PriceCents: 3900},
}

func findProduct(id string) (Product, bool) {
	for _, product := range products {
		if product.ID == id {
			return product, true
		}
	}
	return Product{}, false
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /products", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"products": products})
	})
	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		product, ok := findProduct(r.PathValue("id"))
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
			return
		}
		writeJSON(w, http.StatusOK, product)
	})
	return mux
}
