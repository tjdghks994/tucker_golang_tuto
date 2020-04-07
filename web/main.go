package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type fooHandler struct {
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad request : ", err)
		return
	}

	user.CreatAt = time.Now()
	data, _ := json.Marshal(&user)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

// User Account
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatAt   time.Time `json:"creat_at"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}

		fmt.Fprint(w, "Hello ", name)
	})

	mux.Handle("/foo", &fooHandler{})

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println(err)
	}
}
