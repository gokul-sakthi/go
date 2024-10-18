package main

import (
	"fmt"
	"strings"
)

type Array []string

func NewArray() *Array {
	return &Array{}
}

func (a Array) Append(item string) *Array {
	a = append(a, item)
	return &a
}

func (a Array) At(ind int) (string, error) {
	if ind < 0 || ind >= len(a) {
		return "", fmt.Errorf("Array index (%v) out of bounds", ind)
	}

	return a[ind], nil
}

func (a Array) Size() int {
	return len(a)
}

func (a Array) ToString() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for ind, item := range a {
		builder.WriteString(item)

		if ind != len(a)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")
	return builder.String()
}
