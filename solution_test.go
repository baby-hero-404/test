// Package main contains unit tests for the solution package.
package main

import "testing"

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "existing numbers",
			arr:      []int{1, 2, 3, 2, 4, 2, 5},
			target:   2,
			expected: 3,
		},
		{
			name:     "missing numbers",
			arr:      []int{1, 2, 3},
			target:   4,
			expected: 0,
		},
		{
			name:     "empty array",
			arr:      []int{},
			target:   1,
			expected: 0,
		},
		{
			name:     "all target elements",
			arr:      []int{5, 5, 5, 5},
			target:   5,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountOccurrences(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("CountOccurrences(%v, %d) = %d; want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
