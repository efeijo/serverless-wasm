package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

type AddOp struct {
	A int `json:"A"`
	B int `json:"B"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	addOps := AddOp{}

	err = json.Unmarshal(res, &addOps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "The result for your Operation is %d\n", addOps.A+addOps.B)
	w.Header().Set("Content-Type", "text/plain")
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintln(w, "Hello Fermyon!")
			w.Header().Set("Content-Type", "text/plain")
		case http.MethodPost:
			Add(w, r)
		}
	})
}

func main() {}
