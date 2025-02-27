package caching_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	cache "caching"
)


func TestCache(t *testing.T) {
	c := cache.New[int, string]()

	c.Upsert(5, "funf")

	v, found := c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "funf", v)
}