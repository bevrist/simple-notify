package main

import (
	"fmt"
	"testing"

	"github.com/gorilla/mux"
)

// Test_setupRouterNoDuplicatePaths verifies that the router does not contain duplicate route paths
func Test_setupRouterNoDuplicatePaths(t *testing.T) {
	r := setupRouter()
	// iterate through routes and ensure no duplicate paths
	paths := make(map[string]bool)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			t.Errorf("Error getting path: %v", err)
			t.FailNow()
		}
		if paths[path] {
			t.Errorf("Duplicate path: %v", path)
			t.FailNow()
		}
		paths[path] = true
		return nil
	})
}

// Test_setupRouterNoDuplicateHandlers verifies that no routes use the same handler (or use nil handlers)
func Test_setupRouterNoDuplicateHandlers(t *testing.T) {
	r := setupRouter()
	// iterate through routes and ensure no duplicate handlers
	handlers := make(map[string]string)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		handler := route.GetHandler()
		handName := fmt.Sprint(handler)
		path, _ := route.GetPathTemplate()
		if handName == "<nil>" {
			t.Errorf("Handler is nil for path: %v", path)
			t.FailNow()
		}
		// ignore if handler is bound to hostname
		host, _ := route.GetHostTemplate()
		if host != "" {
			return nil
		}
		if handlers[handName] != "" {
			t.Errorf("Duplicate handler found on paths: %v, %v", path, handlers[handName])
			t.FailNow()
		}
		handlers[handName] = path
		return nil
	})
}

// Test_setupRouterNoDuplicateHostHandlers verifies that no hostname-bound routes use the same handler (or use nil handlers)
func Test_setupRouterNoDuplicateHostHandlers(t *testing.T) {
	r := setupRouter()
	// iterate through routes and ensure no duplicate handlers
	handlers := make(map[string]string)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		handler := route.GetHandler()
		handName := fmt.Sprint(handler)
		path, _ := route.GetPathTemplate()
		if handName == "<nil>" {
			t.Errorf("Handler is nil for path: %v", path)
			t.FailNow()
		}
		// ignore if handler is NOT bound to hostname
		host, _ := route.GetHostTemplate()
		if host == "" {
			return nil
		}
		path = host + path
		if handlers[handName] != "" {
			t.Errorf("Duplicate handler found on host paths: %v, %v", path, handlers[handName])
			t.FailNow()
		}
		handlers[handName] = path
		return nil
	})
}

// Test_setupRouterEnsureMethods verifies that all routes use at least one method
func Test_setupRouterEnsureMethods(t *testing.T) {
	r := setupRouter()
	// iterate through routes and ensure all implement a method
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		_, err := route.GetMethods()
		path, _ := route.GetPathTemplate()
		if err != nil {
			t.Errorf("Error getting methods for path: %v, %v", path, err)
			t.FailNow()
		}
		return nil
	})
}
