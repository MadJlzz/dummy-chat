package structure

import "testing"

func TestAdd(t *testing.T) {
	s := Set{}
	s.Add("Frodo")
	if len(s) != 1 {
		t.Errorf("Add operation failed, got: %d, want: %d.", len(s), 1)
	}
	s.Add("Frodo")
	if len(s) != 1 {
		t.Errorf("Add operation should not create duplicates, got: %d, want: %d.", len(s), 1)
	}
}

func TestRemove(t *testing.T) {
	s := Set{}
	s.Add("Sauron")
	s.Remove("Sauron")
	if len(s) != 0 {
		t.Errorf("Remove operation failed, got: %d, want: %d.", len(s), 0)
	}
}

func TestSize(t *testing.T) {
	s := Set{}
	s.Add("Gandalf")
	size := s.Size()
	if size != len(s) {
		t.Errorf("Wrong Set size, got: %d, want: %d.", size, 1)
	}
}

func TestHas(t *testing.T) {
	s := Set{}
	if ok := s.Has("Saroumane"); ok {
		t.Errorf("Key should not exist, got: %t, want: %t.", ok, false)
	}
	s.Add("Saroumane")
	if ok := s.Has("Saroumane"); !ok {
		t.Errorf("Key should exist now, got: %t, want: %t.", ok, true)
	}
}
