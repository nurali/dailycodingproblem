package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSumUp(t *testing.T) {
	tables := []struct {
		numbers []int
		no      int
		want    bool
	}{
		{[]int{10, 15, 3, 7}, 17, true},
		{[]int{10, 15, 3, 7}, 16, false},
		{[]int{10, 15, 3, 7}, 16, false},
		{[]int{17, 15, 3, 7}, 17, false},
		{[]int{17, 15, 0, 7}, 17, true},
	}

	for _, table := range tables {
		got := CanSumUp(table.numbers, table.no)
		assert.Equal(t, table.want, got)
	}
}
