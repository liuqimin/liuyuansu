package main

import (
	"fmt"
	"log"
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

type MiddlewareList []MiddlewareFunc

func (mList MiddlewareList) Use(handler http.Handler) http.Handler {
	for i := len(mList) - 1; i >= 0; i-- {
		handler = mList[i](handler)
	}
	return handler
}

func LDAPAuth() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("LDAPAuth Received request: %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}

func SAMlAuth() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("SAMlAuth Received request: %s %s\n", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	MiddlewareListNow := MiddlewareList{
		LDAPAuth(),
		SAMlAuth(),
	}
	finalHandler := MiddlewareListNow.Use(http.HandlerFunc(HelloHandler))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
