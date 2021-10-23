package store

import "testing"

func TestRepoPutAndGet(t *testing.T) {
	repo := NewRepository()

	cases := []struct {
		in   string
		want string
	}{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
	}
	for _, c := range cases {
		repo.Put(c.in, c.want)

		got, found := repo.Get(c.in)
		if !found {
			t.Errorf("repo.Get(%q) not found", c.in)
		}

		if got != c.want {
			t.Errorf("repo.Get(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
