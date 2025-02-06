package calculator

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		total     int
		packSizes []int
		expected  map[int]int
	}{
		{1, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},
		{250, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},
		{251, []int{250, 500, 1000, 2000, 5000}, map[int]int{500: 1}},
		{501, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1, 500: 1}},
		{12001, []int{250, 500, 1000, 2000, 5000}, map[int]int{5000: 2, 2000: 1, 250: 1}},
	}

	for _, test := range tests {
		result := CalculatePacks(test.total, test.packSizes)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CalculatePacks(%d, %v) = %v; expected %v", test.total, test.packSizes, result, test.expected)
		}
	}
}
