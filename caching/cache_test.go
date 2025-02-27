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

	c.Upsert(5, "pum")

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	v, found = c.Read(5)
	assert.False(t, found)
	assert.Equal(t, "", v)

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)
}