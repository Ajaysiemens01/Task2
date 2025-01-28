package set

import (
	"fmt"
	"testing"
)

// Helper function to check equality of two slices
func equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	valCount := make(map[interface{}]int)
	for _, val := range a {
		valCount[val]++
	}
	for _, val := range b {
		if valCount[val] == 0 {
			return false
		}
		valCount[val]--
	}
	return true
}

var s Set
// TestUnion tests the Union function of the set package
func TestUnion(t *testing.T) {
	slice1 := []interface{}{1, 2, 3}
	slice2 := []interface{}{3, 4, 5}

	expected := []interface{}{1, 2, 3, 4, 5}
	result, err := s.Union(slice1, slice2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !equal(result, expected) {
		t.Errorf("TestUnion failed: Expected %v, got %v", expected, result)
	} else {
		fmt.Println("TestUnion PASS")
	}
}

// TestIntersection tests the Intersection function of the set package
func TestIntersection(t *testing.T) {
	slice1 := []interface{}{1, 2, 3}
	slice2 := []interface{}{3, 4, 5}

	expected := []interface{}{3}
	result, err := s.Intersection(slice1, slice2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !equal(result, expected) {
		t.Errorf("TestIntersection failed: Expected %v, got %v", expected, result)
	} else {
		fmt.Println("TestIntersection PASS")
	}
}

// TestIsEqual tests the IsEqual function of the set package
func TestIsEqual(t *testing.T) {
	slice1 := []interface{}{1, 2, 3}
	slice2 := []interface{}{1, 2.0, 3}

	expected := false
	result, err := s.IsEqual(slice1, slice2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("TestIsEqual failed: Expected %v, got %v", expected, result)
	} else {
		fmt.Println("TestIsEqual PASS")
	}
}

// TestIsSubset tests the IsSubset function of the set package
func TestIsSubset(t *testing.T) {
	slice1 := []interface{}{1, '2'}
	slice2 := []interface{}{1, '2', 3}

	expected := true
	result, err := s.IsSubset(slice1, slice2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("TestIsSubset failed: Expected %v, got %v", expected, result)
	} else {
		fmt.Println("TestIsSubset PASS")
	}
}

// TestIsProperSubset tests the IsProperSubset function of the set package
func TestIsProperSubset(t *testing.T) {
	slice1 := []interface{}{1.1, "ok"}
	slice2 := []interface{}{1.1, 2, "ok"}

	expected := true
	result, err := s.IsProperSubset(slice1, slice2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("TestIsProperSubset failed: Expected %v, got %v", expected, result)
	} else {
		fmt.Println("TestIsProperSubset PASS")
	}
}

// TestComplement tests the Complement function of the set package
func TestComplement(t *testing.T) {
	union := []interface{}{1, 2, 3, 4}
	slice := []interface{}{2, 3}

	expected := []interface{}{1, 4}
	result, err := s.Complement(union, slice)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !equal(result, expected) {
		t.Errorf("TestComplement failed: Expected %v, got %v", expected, result)
	} else {
		fmt.Println("TestComplement PASS")
	}
}

