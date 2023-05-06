package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routes_exist(t *testing.T) {
	testApp := &Config{}

	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	routes := []string{"/authenticate"}

	for _, route := range routes {
		routesExist(t, chiRoutes, route)
	}
}

func routesExist(t *testing.T, router chi.Routes, r string) {
	found := false

	_ = chi.Walk(router, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == r {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find %s in registered routes", r)
	}
}
