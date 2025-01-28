package set

import (
	"errors"
)

type Set struct{}

// Union returns the union of two slices.
func (s Set) Union(slice1, slice2 []interface{}) ([]interface{}, error) {
	if slice1 == nil || slice2 == nil {
		return nil, errors.New("input slices cannot be nil")
	}
	result := []interface{}{}
	valCount := make(map[interface{}]bool)
	for _, val := range slice1 {
		if !valCount[val] {
			valCount[val] = true
			result = append(result, val)
		}
	}
	for _, val := range slice2 {
		if !valCount[val] {
			valCount[val] = true
			result = append(result, val)
		}
	}
	return result, nil
}

// Intersection returns the intersection of two slices.
func (s Set) Intersection(slice1, slice2 []interface{}) ([]interface{}, error) {
	if slice1 == nil || slice2 == nil {
		return nil, errors.New("input slices cannot be nil")
	}
	result := []interface{}{}
	valCount := make(map[interface{}]bool)
	for _, val := range slice1 {
		valCount[val] = true
	}
	for _, val := range slice2 {
		if valCount[val] {
			result = append(result, val)
		}
	}
	return result, nil
}

// IsEqual checks if two slices are equal.
func (s Set) IsEqual(slice1, slice2 []interface{}) (bool, error) {
	if slice1 == nil || slice2 == nil {
		return false, errors.New("input slices cannot be nil")
	}
	if len(slice1) != len(slice2) {
		return false, nil
	}
	valCount := make(map[interface{}]int)
	for _, val := range slice1 {
		valCount[val]++
	}
	for _, val := range slice2 {
		if valCount[val] == 0 {
			return false, nil
		}
		valCount[val]--
	}
	return true, nil
}

// IsSubset checks if slice1 is a subset of slice2.
func (s Set) IsSubset(slice1, slice2 []interface{}) (bool, error) {
	if slice1 == nil || slice2 == nil {
		return false, errors.New("input slices are empty")
	}
	valCount := make(map[interface{}]bool)
	for _, val := range slice2 {
		valCount[val] = true
	}
	for _, val:= range slice1 {
		if !valCount[val] {
			return false, nil
		}
	}
	return true, nil
}

// IsProperSubset checks if slice1 is a proper subset of slice2.
func (s Set) IsProperSubset(slice1, slice2 []interface{}) (bool, error) {
	isSubset, err := s.IsSubset(slice1, slice2)
	if err != nil {
		return false, err
	}
	return isSubset && len(slice1) < len(slice2), nil
}

// Complement returns the complement of slice in the union.
func (s Set) Complement(union, slice []interface{}) ([]interface{}, error) {
	if union == nil || slice == nil {
		return nil, errors.New("input slices cannot be nil")
	}
	result := []interface{}{}
	valCount := make(map[interface{}]bool)
	for _, val := range slice {
		valCount[val] = true
	}
	for _, val := range union {
		if !valCount[val] {
			result = append(result, val)
		}
	}
	return result, nil
}


