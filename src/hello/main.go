package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Member struct {
	Firstname string `json:"firstname"`
	Code      int    `json:"code"`
	Phone     string `json:"phone"`
}

func main() {
	var m []Member

	a := Member{
		Firstname: "John",
		Code:      123,
		Phone:     "0979232669",
	}

	b := Member{
		Firstname: "Steve",
		Code:      456,
		Phone:     "0992345667",
	}

	m = append(m, a)
	m = append(m, b)

	//fmt.Println(m)

	routeRequest(m)
}

func routeRequest(m []Member) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/member", member(m))
	http.ListenAndServe(":3333", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my homepage")
	fmt.Fprint(w, "Hello my world")
}

func member(m []Member) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			/* m := []struct {
				Firstname string
				Code      int
				Phone     string
			}{
				{
					Firstname: "John",
					Code:      123,
					Phone:     "0979232669",
				},
				{
					Firstname: "Steve",
					Code:      456,
					Phone:     "0979225668",
				},
			} */

			js, _ := json.Marshal(m)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else {
			fmt.Fprint(w, "only GET method")
		}
	}
}

/* func memberPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		m := []struct {
			Firstname string
			Code      int
			Phone     string
		}{
			{
				Firstname: "John",
				Code:      123,
				Phone:     "0979232669",
			},
			{
				Firstname: "Steve",
				Code:      456,
				Phone:     "0979225668",
			},
		}

		js, _ := json.Marshal(m)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		fmt.Fprint(w, "only GET method")
	}
} */
