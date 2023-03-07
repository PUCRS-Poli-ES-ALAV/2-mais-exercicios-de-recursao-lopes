package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFatorial(t *testing.T) {

	testCases := []struct {
		input    int
		output   int
		hasError bool
	}{
		{input: 0, output: 1, hasError: false},
		{input: 1, output: 1, hasError: false},
		{input: 2, output: 2, hasError: false},
		{input: 3, output: 6, hasError: false},
		{input: 4, output: 24, hasError: false},
		{input: 5, output: 120, hasError: false},
		{input: -1, output: 0, hasError: true},
	}

	for _, tc := range testCases {
		output, err := Fatorial(tc.input)
		if tc.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.output, output)
	}
}

func TestSomatorio(t *testing.T) {

	testCases := []struct {
		input  int
		output int
	}{
		{input: 0, output: 0},
		{input: 1, output: 1},
		{input: 2, output: 3},
		{input: 3, output: 6},
		{input: 4, output: 10},
		{input: 5, output: 15},
	}

	for _, tc := range testCases {
		output := Somatorio(tc.input)
		assert.Equal(t, tc.output, output)
	}
}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{input: 0, output: 0},
		{input: 1, output: 1},
		{input: 2, output: 1},
		{input: 3, output: 2},
		{input: 4, output: 3},
		{input: 5, output: 5},
		{input: 6, output: 8},
		{input: 7, output: 13},
		{input: 8, output: 21},
	}

	for _, tc := range testCases {
		output := Fibonacci(tc.input)
		assert.Equal(t, tc.output, output)
	}
}

func TestSomatorioIntervalo(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{input: []int{0, 0}, output: 0},
		{input: []int{0, 1}, output: 1},
		{input: []int{1, 0}, output: 1},
		{input: []int{1, 1}, output: 1},
		{input: []int{1, 2}, output: 3},
		{input: []int{1, 10}, output: 55},
		{input: []int{10, 1}, output: 55},
	}

	for _, tc := range testCases {
		output := SomatorioIntervalo(tc.input[0], tc.input[1])
		assert.Equal(t, tc.output, output)
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{input: "a", output: true},
		{input: "aa", output: true},
		{input: "ab", output: false},
		{input: "aba", output: true},
		{input: "abba", output: true},
		{input: "abcba", output: true},
		{input: "", output: true},
	}

	for _, tc := range testCases {
		output := IsPalindrome(tc.input)
		assert.Equal(t, tc.output, output, fmt.Sprintf("input: %s output: %t expected: %t", tc.input, tc.output, tc.output))
	}
}

func TestConvIntToBin(t *testing.T) {
	testCases := []struct {
		input  int
		output string
	}{
		{input: 0, output: "0"},
		{input: 1, output: "1"},
		{input: 2, output: "10"},
		{input: 3, output: "11"},
		{input: 4, output: "100"},
		{input: 5, output: "101"},
		{input: 6, output: "110"},
		{input: 7, output: "111"},
		{input: 8, output: "1000"},
	}

	for _, tc := range testCases {
		output := ConvIntToBin(tc.input)
		assert.Equal(t, tc.output, output, fmt.Sprintf("input: %d output: %s expected: %s", tc.input, tc.output, tc.output))
	}
}

func TestSomaLista(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{input: []int{1, 2, 3}, output: 6},
		{input: []int{1, 2, 3, 4}, output: 10},
		{input: []int{1, 2, 3, 4, 5}, output: 15},
		{input: []int{1}, output: 1},
		{input: []int{}, output: 0},
	}
	for _, tc := range testCases {
		output := SomaLista(tc.input)
		assert.Equal(t, tc.output, output)
	}
}
