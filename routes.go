package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported 2"))
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Failed to read body"))
		return
	}

	payload := &LoginPayload{}
	err = json.Unmarshal(bodyBytes, payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := s.keycloakClient.Login(r.Context(), payload.Username, payload.Password)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	rol, err := s.keycloakClient.DecodeJWT(r.Context(), res.AccessToken)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	/*
		var tokenString = res.AccessToken
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			fmt.Println(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["realm_access"])
		} else {
			fmt.Println(err)
		}
	*/

	/*
		rol, err := s.keycloakClient.DecodeJWT(r.Context(), res.AccessToken)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
	*/

	/*
		_, err = s.keycloakClient.GetRoleMappingByUserID(r.Context(), res.AccessToken)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
	*/

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(rol)
}

/* func (s *APIServer) handleRoleMapping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported 2"))
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Failed to read body"))
		return
	}

	payload := &IntrospectPayload{}
	err = json.Unmarshal(bodyBytes, payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

 	res, err := s.keycloakClient.Login(r.Context(), payload.AccessToken)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

*/
