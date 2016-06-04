package hello

import "testing"

func TestHelloFunc(t *testing.T) {
	hel := HelloFunc()

	if hel != "hello" {
		t.Fatalf("expected hello")
	}
	
}

