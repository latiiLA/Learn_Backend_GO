package main

import (
	"testing"
)

func TestAverageComputer(t *testing.T) {
	tests := []struct {
		grades   []float64
		expected float64
	}{
		{[]float64{90, 80, 70}, 80.0},
		{[]float64{100, 100, 100}, 100.0},
		{[]float64{50, 60, 70, 80}, 65.0},
		{[]float64{}, 0.0},
		{[]float64{100}, 100.0},
		{[]float64{-1, -2, -30}, -11.0},
	}

	for _, tt := range tests {
		t.Run("testing average", func(t *testing.T) {
			got := average_computer(tt.grades)
			if got != tt.expected {
				t.Errorf("average_computer(%v) = %v; want %v", tt.grades, got, tt.expected)
			}
		})
	}
}
