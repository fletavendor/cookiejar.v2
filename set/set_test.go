// CookieJar - A contestant's algorithm toolbox
// Copyright 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The toolbox is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)
package set

import (
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	// Fill the set with the data and verify that they're all set
	set := New()
	for val := range data {
		set.Insert(val)
	}
	for val := range data {
		if !set.Exists(val) {
			t.Errorf("failed to locate element in set: %v in %v", val, set)
		}
	}
	// Remove a few elements and ensure they're out
	rems := data[:1024]
	for val := range rems {
		set.Remove(val)
	}
	for val := range rems {
		if set.Exists(val) {
			t.Errorf("element exists after remove: %v in %v", val, set)
		}
	}
	// Calcualte the sum of the remainder and verify
	sumSet := int64(0)
	set.Do(func(val interface{}) {
		sumSet += int64(val.(int))
	})
	sumDat := int64(0)
	for val := range data {
		sumDat += int64(val)
	}
	for val := range rems {
		sumDat -= int64(val)
	}
	if sumSet != sumDat {
		t.Errorf("iteration sum mismatch: have %v, want %v", sumSet, sumDat)
	}
	// Clear the set and ensure nothing's left
	set.Reset()
	for val := range data {
		if set.Exists(val) {
			t.Errorf("element exists after reset: %v in %v", val, set)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Execute the benchmark
	b.ResetTimer()
	set := New()
	for val := range data {
		set.Insert(val)
	}
}

func BenchmarkRemove(b *testing.B) {
	// Create some initial data and fill the set
	data := rand.Perm(b.N)
	set := New()
	for val := range data {
		set.Insert(val)
	}
	// Execute the benchmark (different order)
	rems := rand.Perm(b.N)
	b.ResetTimer()
	for val := range rems {
		set.Remove(val)
	}
}

func BenchmarkDo(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Fill the set with it
	set := New()
	for val := range data {
		set.Insert(val)
	}
	// Execute the benchmark
	b.ResetTimer()
	set.Do(func(val interface{}) {
		// Do nothing
	})
}
