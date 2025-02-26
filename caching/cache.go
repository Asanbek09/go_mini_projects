package main

type Cache[K comparable, V any] struct {
	data map[K]V
}