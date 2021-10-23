package store

import "testing"

func TestServicePutAndGet(t *testing.T) {
	service := NewService(NewRepository())

	cases := []struct {
		in   string
		want string
	}{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
	}
	for _, c := range cases {
		service.Put(c.in, c.want)

		got, found := service.Get(c.in)
		if !found {
			t.Errorf("service.Get(%q) not found", c.in)
		}

		if got != c.want {
			t.Errorf("service.Get(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
