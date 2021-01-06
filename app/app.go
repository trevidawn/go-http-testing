package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/foo", FooHandler)
	mux.Handle("/bar", &BarHandler{})

	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
	w.WriteHeader(http.StatusOK)
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello %s!", name)
	w.WriteHeader(http.StatusOK)
}

type BarHandler struct{}

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (handler *BarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}
