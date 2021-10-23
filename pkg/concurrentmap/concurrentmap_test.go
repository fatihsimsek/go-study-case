package concurrentmap

import (
	"testing"
)

func TestPutGet(t *testing.T) {
	m := New()

	cases := []struct {
		in   string
		want interface{}
	}{
		{"key1", 432},
		{"key2", "deneme"},
		{"key1", "deneme2"},
	}
	for _, c := range cases {
		m.Put(c.in, c.want)

		got, found := m.Get(c.in)
		if !found {
			t.Errorf("m.Get(%q) not found", c.in)
		}

		if got != c.want {
			t.Errorf("m.Get(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestContains(t *testing.T) {
	m := New()

	if m.Contains("key1") {
		t.Errorf("m.Contains(key1) == true")
	}

	m.Put("key1", 1)

	if !m.Contains("key1") {
		t.Errorf("m.Contains(key1) == false")
	}
}

func TestRemove(t *testing.T) {
	m := New()

	m.Put("key1", 1)

	if found := m.Remove("key1"); !found {
		t.Errorf("m.Remove(key1) == false")
	}
}
