package linked_list

import (
	"testing"
)

var list List

func TestList_IsEmpty(t *testing.T) {
	got := list.IsEmpty()
	want := true

	if got != want {
		t.Errorf("Expected %t got %t", want, got)
	}
}

func TestList_Size(t *testing.T) {
	got := list.Size()
	want := 0

	if got != want {
		t.Errorf("Expected %d got %d", want, got)
	}
}

func TestList_Append(t *testing.T) {
	list.Append("Robinson Crusoe")

	size := list.Size()
	expected := 1
	if size != expected {
		t.Errorf("Expected %d got %d", expected, size)
	}

	list.Append("Data Structures")
	list.Append("Algoritms")

	size = list.Size()
	expected = 3
	if size != expected {
		t.Errorf("Expected %d got %d", expected, size)
	}

	list.Append("Clean Code")
	list.Append("Problem Solving with Algorithms and Data Structures")

	size = list.Size()
	expected = 5
	if size != expected {
		t.Errorf("Expected %d got %d", expected, size)
	}
}

func TestList_AppendAt(t *testing.T) {
	err := list.AppendAt(100, "Sandman")

	if err == nil {
		t.Errorf("Waiting for the error got %s", err)
	}

	err = list.AppendAt(3, "Sandman")
	if err != nil {
		t.Errorf("%s", err)
	}

}
