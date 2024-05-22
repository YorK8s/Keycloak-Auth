package main

import (
	"PSKE-API-AUTH/keycloak"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr           string
	client         *Client
	keycloakClient keycloak.Keycloak
}

func NewAPIServer(addr string) *APIServer {
	httpClient := &http.Client{}
	return &APIServer{
		addr: addr,
		client: &Client{
			httpClient: httpClient,
		},
		keycloakClient: *keycloak.InitKeycloak("http://localhost:8181", "auth", "clientSecret", "realm"),
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet)
	router.HandleFunc("/login", s.handleLogin).Methods(http.MethodPost)
	/*	router.HandleFunc("/introspect", s.handleRoleMapping).Methods(http.MethodPost) */

	fmt.Printf("Server starting on address %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
