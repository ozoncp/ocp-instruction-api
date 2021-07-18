package utils

import (
	//"fmt"
	"reflect"
	"testing"
)

func TestBatchSlice(t *testing.T) {
	_, err := BatchIntSlice(make([]int, 0), 10)
	if err == nil {
		t.Errorf("expected error, got no error")
	}

	_, err = BatchIntSlice(make([]int, 10), 0)
	if err == nil {
		t.Errorf("expected error, got no error")
	}

	ret, err := BatchIntSlice([]int{10, 65, 23, 52, 567, 43, 76, 12, 45, 15}, 7)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	exptd := [][]int{{10, 65, 23, 52, 567, 43, 76}, {12, 45, 15}}
	// slow! but simple...
	if !reflect.DeepEqual(exptd, ret) {
		t.Errorf("expected %v, got %v", exptd, ret)
	}
}

func TestSwapMap(t *testing.T) {
	_, err := SwapIntMap(map[int]int{15: 20, 34: 20, 2: 22, 94: 9, 83: 13})
	if err == nil {
		t.Errorf("expected error, got no error")
	}

	res, err := SwapIntMap(map[int]int{15: 20, 34: 21, 2: 22, 94: 9, 83: 13})
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	expectedMap := map[int]int{20: 15, 21: 34, 22: 2, 9: 94, 13: 83}

	if len(res) != len(expectedMap) {
		t.Errorf("expected %v, got %v", expectedMap, res)
	}

	for k := range expectedMap {
		if val, ok := res[k]; !ok || val != expectedMap[k] {
			t.Errorf("expected %v, got %v", expectedMap, res)
			break
		}
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFilterSlice(t *testing.T) {
	res := FilterIntSlice([]int{1, 2, 3}, 4, 5, 6)
	if len(res) != 3 {
		t.Errorf("expected %v, got %v", []int{}, res)
	}
	if !sliceEqual(res, []int{1, 2, 3}) {
		t.Errorf("expected %v, got %v", []int{}, res)
	}

	res = FilterIntSlice([]int{1, 2, 3}, 3, 5, 2)
	if len(res) != 1 {
		t.Errorf("expected %v, got %v", []int{}, res)
	}
	if !sliceEqual(res, []int{1}) {
		t.Errorf("expected %v, got %v", []int{}, res)
	}

	res = FilterIntSlice([]int{1, 2, 3}, 3, 1, 2)
	if len(res) != 0 {
		t.Errorf("expected %v, got %v", []int{}, res)
	}
}
