package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {

	inputArgs := []int{0, -35, 10}

	want := [][]int{
		{},
		{},
		make([]int, 10),
	}

	for i, arg := range inputArgs {
		if len(generateRandomElements(arg)) != len(want[i]) {
			t.Error(i, ":", len(generateRandomElements(arg)), "!=", want[i])
		}
	}
}

func TestMaximum(t *testing.T) {

	inputArgs := [][]int{
		{},
		{3242},
		{1, 5, 3, 9, 2},
	}

	want := []int{0, 3242, 9}

	for i, arg := range inputArgs {
		if maximum(arg) != want[i] {
			t.Error(i, ":", maximum(arg), "!=", want[i])
		}
	}
}

func TestMaxChunks(t *testing.T) {

	inputArgs := [][]int{
		{},
		{3242},
		{1, 5, 3, 9, 2},
	}

	want := []int{0, 3242, 9}

	for i, arg := range inputArgs {
		if maxChunks(arg) != want[i] {
			t.Error(i, ":", maxChunks(arg), "!=", want[i])
		}
	}
}
