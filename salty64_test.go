package salty64

import (
	"strconv"
	"testing"
)

func TestNewShaker(t *testing.T) {
	s, err := NewShaker("test", 1)

	if err != nil {
		t.Fatalf("NewShaker returned error: %s\n", err)
	}

	if s.Salt != "test" {
		t.Fatalf("NewShaker did not apply salt correctly: %+v\n", s)
	}

	if s.Offset != 1 {
		t.Fatalf("NewShaker did not apply offset correctly: %+v\n", s)
	}

	s, err = NewShaker("test", 5)
	if err == nil {
		t.Fatal("NewShaker did not validate correctly\n")
	}

	s, err = NewShaker("test", -1)
	if err == nil {
		t.Fatal("NewShaker did not validate correctly\n")
	}
}

func TestEncode(t *testing.T) {
	s, err := NewShaker("test", 1)
	if err != nil {
		t.Fatalf("NewShaker returned error: %s\n", err)
	}

	h1, err := s.Encode("hello world")
	if err != nil {
		t.Fatalf("Encode returned error: %s\n", err)
	}

	h2, err := s.Encode("hello world")
	if err != nil {
		t.Fatalf("Encode returned error: %s\n", err)
	}

	if len(h1) <= 0 {
		t.Fatalf("Failed to encode")
	}

	if h1 != h2 {
		t.Fatalf("Mismatched encodings on same string")
	}
}

func TestDecode(t *testing.T) {
	s, err := NewShaker("test", 1)
	if err != nil {
		t.Fatalf("NewShaker returned error: %s\n", err)
	}

	enc, err := s.Encode("hello world")
	if err != nil {
		t.Fatalf("Encode returned error: %s\n", err)
	}

	dec, err := s.Decode(enc)
	if err != nil {
		t.Fatalf("Decode returned error: %s\n", err)
	}

	if dec != "hello world" {
		t.Fatalf("Decoded to unexpected string: %q\n", dec)
	}
}

func BenchmarkEncode(b *testing.B) {
	s, err := NewShaker("test", 1)
	if err != nil {
		b.Fatalf("NewShaker returned error: %s\n", err)
	}

	for n := 0; n < b.N; n++ {
		_, err := s.Encode("test run @ " + strconv.Itoa(n))
		if err != nil {
			b.Fatalf("Encode returned error: %s\n", err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	s, err := NewShaker("test", 1)
	if err != nil {
		b.Fatalf("NewShaker returned error: %s\n", err)
	}

	for n := 0; n < b.N; n++ {
		enc, err := s.Encode("test run @ " + strconv.Itoa(n))
		if err != nil {
			b.Fatalf("Encode returned error: %s\n", err)
		}
		_, err = s.Decode(enc)
		if err != nil {
			b.Fatalf("Decode returned error: %s\n", err)
		}
	}
}
