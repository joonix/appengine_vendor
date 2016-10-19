package main

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func handlerWithContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}

func init() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		handlerWithContext(ctx, w, r)
	}))
}
