package hello

import "testing"

func TestHelloFunc(t *testing.T) {
	hel := HelloFunc()

	if hel != "hello" {
		t.Fatalf("expected hello")
	}
	
}

func TestTimesTwo(t *testing.T) {
	two := TimesTwo(4)

	if two != 8 {
		t.Fatalf("expected 8")
	}
}
