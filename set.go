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
	m := make(map[interface{}]bool)
	for _, v := range slice1 {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	for _, v := range slice2 {
		if !m[v] {
			m[v] = true
			result = append(result, v)
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
	m := make(map[interface{}]bool)
	for _, v := range slice1 {
		m[v] = true
	}
	for _, v := range slice2 {
		if m[v] {
			result = append(result, v)
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
	m := make(map[interface{}]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		if m[v] == 0 {
			return false, nil
		}
		m[v]--
	}
	return true, nil
}

// IsSubset checks if slice1 is a subset of slice2.
func (s Set) IsSubset(slice1, slice2 []interface{}) (bool, error) {
	if slice1 == nil || slice2 == nil {
		return false, errors.New("input slices are empty")
	}
	m := make(map[interface{}]bool)
	for _, v := range slice2 {
		m[v] = true
	}
	for _, v := range slice1 {
		if !m[v] {
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
	m := make(map[interface{}]bool)
	for _, v := range slice {
		m[v] = true
	}
	for _, v := range union {
		if !m[v] {
			result = append(result, v)
		}
	}
	return result, nil
}


