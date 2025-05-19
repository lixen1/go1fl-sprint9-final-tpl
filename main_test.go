package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {

	t.Run("Zero size", func(t *testing.T) {
		res := generateRandomElements(0)
		if len(res) != 0 {
			t.Errorf("Want empty slice, have %d", len(res))
		}
	})

	t.Run("Negative size", func(t *testing.T) {
		res := generateRandomElements(-35)
		if len(res) != 0 {
			t.Errorf("Want empty slice, have %d", len(res))
		}
	})
	t.Run("Positive size", func(t *testing.T) {
		size := 10
		result := generateRandomElements(size)
		if len(result) != size {
			t.Errorf("Want %d, have %d", size, len(result))
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		result := maximum([]int{})
		if result != 0 {
			t.Errorf("Want zero, have %d", result)
		}
	})

	t.Run("Single element in slice", func(t *testing.T) {
		result := maximum([]int{3242})
		if result != 3242 {
			t.Errorf("Want 3242, have %d", result)
		}
	})

	t.Run("Multiple elements", func(t *testing.T) {
		result := maximum([]int{1, 5, 3, 9, 2})
		if result != 9 {
			t.Errorf("Want max = 9, have %d", result)
		}
	})

	t.Run("All negative elements", func(t *testing.T) {
		result := maximum([]int{-10, -3, -50})
		if result != 0 {
			t.Errorf("Want zero, have %d", result)
		}
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		result := maxChunks([]int{})
		if result != 0 {
			t.Errorf("Want zero, have %d", result)
		}
	})

	t.Run("Single element in slice", func(t *testing.T) {
		result := maxChunks([]int{238})
		if result != 238 {
			t.Errorf("Want 238, have %d", result)
		}
	})

}
