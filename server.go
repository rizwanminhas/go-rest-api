package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

type userHandlers struct {
	store map[string]User
}

func (h *userHandlers) get(w http.ResponseWriter, r *http.Request) {
	users := make([]User, len(h.store))

	i := 0
	for _, user := range h.store {
		users[i] = user
		i++
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newUserHandlers() *userHandlers {
	return &userHandlers{
		store: map[string]User{
			"id1": User{
				Name:  "rizwan",
				Age:   100,
				Email: "rizwan@rizwan.com",
				ID:    "id1",
			},
		},
	}
}

func main() {
	userHandlers := newUserHandlers()
	http.HandleFunc("/user", userHandlers.get)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
