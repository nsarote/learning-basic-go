package main

import "testing"

func TestAddWhenInput1and2ShouldReturn3(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1,2) should be 3")
	}
}

func TestAdd(t *testing.T) {
	t.Run("when input 4 and 5 should return 9", func(t *testing.T) {
		if Add(4, 5) != 9 {

			t.Error("Add(4,5) should be 9")
		}
	})

	t.Run("when input 7 and 8 should return 15", func(t *testing.T) {
		if Add(7, 8) != 15 {

			t.Error("Add(7,8) should be 15")
		}
	})
}
