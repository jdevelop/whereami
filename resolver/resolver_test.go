package resolver

import (
	"testing"
)

func TestResolver(t *testing.T) {
	resolver, err := NewResolver()
	if err != nil {
		t.Fatal(err)
	}
	addr, err := resolver.GetAllAddresses()
	if err != nil {
		t.Fatal(err)
	}
	if len(addr) != 1 {
		t.Fatalf("Expected 1, actual %d", len(addr))
	}
}

func TestRouteResolver(t *testing.T) {
	resolver, err := NewResolver()
	if err != nil {
		t.Fatal(err)
	}

	route, err := resolver.GetDefaultRouteSrc()
	if err != nil {
		t.Fatal(err)
	}

	if route == nil {
		t.Fatalf("Can't find default router %v", route)
	}
}
